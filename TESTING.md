# Testing Guide

## Overview

The admin-api uses Go's standard `testing` package with httptest for handler
and route-level unit tests. **214 tests total** across handlers and routes.

## Quick Commands

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Run handler tests only
go test -v ./internal/handlers/

# Run route RBAC tests only
go test -v ./internal/routes/

# Run specific test
go test -v -run TestGetAllCertifications_Success ./internal/handlers/

# Run all Certification tests
go test -v -run Certification ./internal/handlers/

# Run all permission tests
go test -v -run Permission ./internal/routes/
```

## Test Files

### `internal/handlers/handler_test.go` - 81 tests

| Category | Tests | Coverage |
| -------- | ----- | -------- |
| Certifications | 16 | GetAll, GetByID, Create, Update, Delete + errors |
| Skills | 18 | GetAll, GetByID, Create, Update, Delete + errors |
| Skill Types | 6 | GetAll, GetByID, Create, Update, Delete |
| Work Experience | 18 | GetAll, GetByID, Create, Update, Delete + errors |
| Miniature Projects | 14 | GetAll, GetByID, Create, Update, Delete + errors |
| Miniature Techniques | 2 | GetAll + error |
| Project Associations | 4 | SetTechniques, SetPaints + invalid ID |
| Add Image to Project | 3 | Success, InvalidID, MissingFileID |
| Context Propagation | 1 | Verifies context with sentinel value |
| ID Validation | 1 | Table-driven invalid ID format tests |
| Constructor | 1 | Handler initialization |

### `internal/routes/routes_test.go` - 133 tests

| Category | Tests | Coverage |
| -------- | ----- | -------- |
| Portfolio Routes Forbidden | 31 | All portfolio routes return 403 without permission |
| Portfolio Routes Allowed | 31 | All portfolio routes accessible with correct permission |
| Miniatures Routes Forbidden | 19 | All miniature routes return 403 without permission |
| Miniatures Routes Allowed | 19 | All miniature routes accessible with correct permission |
| Files Routes Forbidden | 1 | DELETE /files/:id returns 403 without permission |
| Files Routes Allowed | 1 | DELETE /files/:id accessible with delete permission |
| Permission Hierarchy | 10 | delete > edit > read > none hierarchy |
| Cross-Resource Permissions | 1 | Resource isolation (profile:delete ≠ experience:read) |
| Multiple Resource Permissions | 8 | Mixed permission levels across resources |
| No Scopes | 1 | Returns 401 Unauthorized when scopes missing |
| Invalid Scopes Format | 1 | Returns 500 when scopes is wrong type |
| Repository Error Propagation | 1 | Handler errors pass through middleware |

## Key Testing Patterns

**Mock Repository**: Function fields allow per-test behavior customization

```go
mockRepo.getAllCertificationsFunc = func(
    ctx context.Context,
) ([]models.Certification, error) {
    return expectedCerts, nil
}
```

**HTTP Testing**: Uses `httptest.ResponseRecorder` with Gin router

```go
w := performRequest(router, "GET", "/certifications/1", nil)
if w.Code != http.StatusOK { ... }
```

**Table-driven tests**: Multiple scenarios with `tests := []struct{...}`

**Test Helpers**: Factory functions for consistent test data

```go
cert := createTestCertification()
skill := createTestSkill()
exp := createTestWorkExperience()
project := createTestMiniatureProject()
technique := createTestMiniatureTechnique()
```

**RBAC Testing**: Scope injection simulates JWT claims

```go
// Inject scopes into Gin context (simulates ValidateToken middleware)
func injectScopes(scopes map[string]string) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("scopes", scopes)
        c.Next()
    }
}

// Test with specific permissions
scopes := map[string]string{common.ResourceProfile: common.LevelRead}
router := setupRouterWithScopes(t, scopes)
```

## Test Categories

### Success Cases

- Returns expected data
- Sets correct HTTP status
- Sets Location header on create

### Error Cases

- Repository errors (500)
- Not found errors (404)
- Invalid ID format (400)
- Validation errors (400)

### RBAC Cases (routes_test.go)

- Permission denied (403) - missing or insufficient permission
- Permission granted - exact match or higher level
- Unauthorized (401) - no scopes in context
- Internal error (500) - invalid scopes format

## Contributing Tests

1. Follow naming: `Test<HandlerName>_<Scenario>`
2. Organize by entity with section markers
3. Use table-driven tests for multiple scenarios
4. Mock only the repository methods needed
5. Verify: `go test -cover ./internal/handlers/`

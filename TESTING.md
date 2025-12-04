# Testing Guide

## Overview

The admin-api uses Go's standard `testing` package with httptest for handler
unit tests.

## Quick Commands

```bash
# Run all tests
go test ./internal/handlers/

# Run with coverage
go test -cover ./internal/handlers/

# Generate coverage report
go test -coverprofile=coverage.out ./internal/handlers/
go tool cover -html=coverage.out -o coverage.html

# Run specific test
go test -v -run TestGetAllCertifications_Success ./internal/handlers/

# Run all Certification tests
go test -v -run Certification ./internal/handlers/

# Run all Skill tests
go test -v -run Skill ./internal/handlers/

# Run all Work Experience tests
go test -v -run WorkExperience ./internal/handlers/

# Run all Miniature tests
go test -v -run Miniature ./internal/handlers/
```

## Test Files

**`handler_test.go`** - 83 tests

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

## Contributing Tests

1. Follow naming: `Test<HandlerName>_<Scenario>`
2. Organize by entity with section markers
3. Use table-driven tests for multiple scenarios
4. Mock only the repository methods needed
5. Verify: `go test -cover ./internal/handlers/`

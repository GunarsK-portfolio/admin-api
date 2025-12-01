# Admin API

![CI](https://github.com/GunarsK-portfolio/admin-api/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/GunarsK-portfolio/admin-api)](https://goreportcard.com/report/github.com/GunarsK-portfolio/admin-api)
[![codecov](https://codecov.io/gh/GunarsK-portfolio/admin-api/graph/badge.svg?token=PISWS6LDK4)](https://codecov.io/gh/GunarsK-portfolio/admin-api)
[![CodeRabbit](https://img.shields.io/coderabbit/prs/github/GunarsK-portfolio/admin-api?label=CodeRabbit&color=2ea44f)](https://coderabbit.ai)

RESTful API for managing portfolio content with authentication.

## Features

- Full CRUD operations for portfolio content
- JWT authentication via auth-service
- Profile, work experience, and certifications management
- Skills and skill types management
- Portfolio projects management
- Miniature painting projects and themes management
- Image deletion (deletes file record associations)
- RESTful API with Swagger documentation
- Protected endpoints with middleware

## Tech Stack

- **Language**: Go 1.25.3
- **Framework**: Gin
- **Database**: PostgreSQL (GORM)
- **Common**: portfolio-common library (shared database utilities, auth middleware)
- **Auth**: JWT validation via auth-service
- **Documentation**: Swagger/OpenAPI

## Prerequisites

- Go 1.25+
- Node.js 22+ and npm 11+
- PostgreSQL (or use Docker Compose)
- auth-service running

## Project Structure

```text
admin-api/
├── cmd/
│   └── api/              # Application entrypoint
├── internal/
│   ├── config/           # Configuration
│   ├── handlers/         # HTTP handlers
│   ├── middleware/       # Custom middleware
│   ├── models/           # Data models
│   ├── repository/       # Data access layer
│   ├── routes/           # Route definitions
│   ├── service/          # Business logic
│   └── storage/          # Storage utilities
└── docs/                 # Swagger documentation
```

## Quick Start

### Using Docker Compose

```bash
docker-compose up -d
```

### Local Development

1. Copy environment file:

```bash
cp .env.example .env
```

1. Update `.env` with your configuration:

```env
PORT=8083
DB_HOST=localhost
DB_PORT=5432
DB_USER=portfolio_admin
DB_PASSWORD=portfolio_admin_dev_pass
DB_NAME=portfolio
AUTH_SERVICE_URL=http://localhost:8084/api/v1
FILES_API_URL=http://localhost:8085/api/v1
```

1. Start infrastructure and auth-service (if not running):

```bash
# From infrastructure directory
docker-compose up -d postgres flyway auth-service
```

1. Run the service:

```bash
go run cmd/api/main.go
```

## Available Commands

Using Task:

```bash
# Development
task dev:swagger         # Generate Swagger documentation
task dev:install-tools   # Install dev tools (golangci-lint, govulncheck, etc.)

# Build and run
task build               # Build binary
task test                # Run tests
task test:coverage       # Run tests with coverage report
task clean               # Clean build artifacts

# Code quality
task format              # Format code with gofmt
task tidy                # Tidy and verify go.mod
task lint                # Run golangci-lint
task vet                 # Run go vet

# Security
task security:scan       # Run gosec security scanner
task security:vuln       # Check for vulnerabilities with govulncheck

# Docker
task docker:build        # Build Docker image
task docker:run          # Run service in Docker container
task docker:stop         # Stop running Docker container
task docker:logs         # View Docker container logs

# CI/CD
task ci:all              # Run all CI checks
```

Using Go directly:

```bash
go run cmd/api/main.go                      # Run
go build -o bin/admin-api cmd/api/main.go   # Build
go test ./...                                # Test
```

## API Endpoints

Base URL: `http://localhost:8083/api/v1`

All endpoints (except health check) require JWT authentication via
`Authorization: Bearer <token>` header.

### Health Check

- `GET /health` - Service health status

### Portfolio Domain

All portfolio endpoints are under `/portfolio` path.

#### Profile

- `GET /portfolio/profile` - Get profile information
- `PUT /portfolio/profile` - Update profile
- `PUT /portfolio/profile/avatar` - Update profile avatar (by file ID)
- `DELETE /portfolio/profile/avatar` - Remove profile avatar
- `PUT /portfolio/profile/resume` - Update profile resume (by file ID)
- `DELETE /portfolio/profile/resume` - Remove profile resume

#### Work Experience

- `GET /portfolio/experience` - List all work experience
- `POST /portfolio/experience` - Create work experience entry
- `GET /portfolio/experience/:id` - Get work experience by ID
- `PUT /portfolio/experience/:id` - Update work experience
- `DELETE /portfolio/experience/:id` - Delete work experience

#### Certifications

- `GET /portfolio/certifications` - List all certifications
- `POST /portfolio/certifications` - Create certification
- `GET /portfolio/certifications/:id` - Get certification by ID
- `PUT /portfolio/certifications/:id` - Update certification
- `DELETE /portfolio/certifications/:id` - Delete certification

#### Skills

- `GET /portfolio/skills` - List all skills
- `POST /portfolio/skills` - Create new skill
- `GET /portfolio/skills/:id` - Get skill by ID
- `PUT /portfolio/skills/:id` - Update skill
- `DELETE /portfolio/skills/:id` - Delete skill

#### Skill Types

- `GET /portfolio/skill-types` - List all skill types
- `POST /portfolio/skill-types` - Create skill type
- `GET /portfolio/skill-types/:id` - Get skill type by ID
- `PUT /portfolio/skill-types/:id` - Update skill type
- `DELETE /portfolio/skill-types/:id` - Delete skill type

#### Portfolio Projects

- `GET /portfolio/projects` - List all portfolio projects
- `POST /portfolio/projects` - Create new portfolio project
- `GET /portfolio/projects/:id` - Get portfolio project by ID
- `PUT /portfolio/projects/:id` - Update portfolio project
- `DELETE /portfolio/projects/:id` - Delete portfolio project

### Miniatures Domain

All miniature endpoints are under `/miniatures` path.

#### Miniature Themes

- `GET /miniatures/themes` - List all miniature themes
- `POST /miniatures/themes` - Create miniature theme
- `GET /miniatures/themes/:id` - Get miniature theme by ID
- `PUT /miniatures/themes/:id` - Update miniature theme
- `DELETE /miniatures/themes/:id` - Delete miniature theme

#### Miniature Projects

- `GET /miniatures/projects` - List all miniature projects
- `POST /miniatures/projects` - Create miniature project
- `GET /miniatures/projects/:id` - Get miniature project by ID
- `PUT /miniatures/projects/:id` - Update miniature project
- `DELETE /miniatures/projects/:id` - Delete miniature project

### Files

Generic file deletion endpoint (works for all file types: avatars,
resumes, project images, miniature images).

- `DELETE /files/:id` - Delete file by ID (removes file record and associations)

## Swagger Documentation

When running, Swagger UI is available at:

- `http://localhost:8083/swagger/index.html`

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8083` |
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | Database user | `portfolio_admin` |
| `DB_PASSWORD` | Database password | `portfolio_admin_dev_pass` |
| `DB_NAME` | Database name | `portfolio` |
| `DB_SSLMODE` | PostgreSQL SSL mode | `disable` |
| `AUTH_SERVICE_URL` | Auth service URL | `http://localhost:8084/api/v1` |
| `FILES_API_URL` | Files API URL (file URLs) | `http://localhost:8085/api/v1` |

## Authentication

This API validates JWT tokens issued by auth-service using the
portfolio-common auth middleware. The middleware:

1. Extracts token from `Authorization: Bearer <token>` header
2. Validates token with auth-service
3. Injects user information into request context

All endpoints except `/health` require authentication.

## Integration

This API is consumed by the admin-web frontend for portfolio content management.

## License

[MIT](LICENSE)

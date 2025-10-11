# Admin API

RESTful API for managing portfolio content with authentication.

## Features

- Full CRUD operations for portfolio content
- JWT authentication via auth-service
- Projects, skills, experience management
- Image upload/management via MinIO/S3
- RESTful API with Swagger documentation
- Protected endpoints with middleware

## Tech Stack

- **Language**: Go 1.25
- **Framework**: Gin
- **Database**: PostgreSQL (GORM)
- **Storage**: MinIO (S3-compatible)
- **Auth**: JWT validation via auth-service
- **Documentation**: Swagger/OpenAPI

## Prerequisites

- Go 1.25+
- PostgreSQL (or use Docker Compose)
- MinIO (or use Docker Compose)
- auth-service running

## Project Structure

```
admin-api/
├── cmd/
│   └── api/              # Application entrypoint
├── internal/
│   ├── config/           # Configuration
│   ├── database/         # Database connection
│   ├── handlers/         # HTTP handlers
│   ├── middleware/       # JWT auth middleware
│   ├── models/           # Data models
│   ├── repository/       # Data access layer
│   ├── service/          # Business logic
│   └── storage/          # S3/MinIO integration
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

2. Update `.env` with your configuration:
```env
PORT=8083
DB_HOST=localhost
DB_PORT=5432
DB_USER=portfolio_user
DB_PASSWORD=portfolio_pass
DB_NAME=portfolio
AUTH_SERVICE_URL=http://localhost:8084
S3_ENDPOINT=http://localhost:9000
S3_ACCESS_KEY=minioadmin
S3_SECRET_KEY=minioadmin
S3_BUCKET=images
S3_USE_SSL=false
```

3. Start infrastructure and auth-service (if not running):
```bash
# From infrastructure directory
docker-compose up -d postgres minio flyway auth-service
```

4. Run the service:
```bash
task run
# or
go run cmd/api/main.go
```

## Available Commands

Using Task:
```bash
task run           # Run the service
task build         # Build binary
task test          # Run tests
task swagger       # Generate Swagger docs
task clean         # Clean build artifacts
task docker-build  # Build Docker image
task docker-run    # Run with docker-compose
task docker-logs   # View logs
```

Using Go directly:
```bash
go run cmd/api/main.go       # Run
go build -o bin/admin-api cmd/api/main.go  # Build
go test ./...                # Test
```

## API Endpoints

Base URL: `http://localhost:8083/api/v1`

All endpoints (except health check) require JWT authentication via `Authorization: Bearer <token>` header.

### Health Check
- `GET /health` - Service health status

### Projects
- `GET /projects` - List all projects
- `GET /projects/:id` - Get project details
- `POST /projects` - Create new project
- `PUT /projects/:id` - Update project
- `DELETE /projects/:id` - Delete project

### Skills
- `GET /skills` - List all skills
- `POST /skills` - Create new skill
- `PUT /skills/:id` - Update skill
- `DELETE /skills/:id` - Delete skill

### Experience
- `GET /experience` - List work experience
- `POST /experience` - Create experience entry
- `PUT /experience/:id` - Update experience
- `DELETE /experience/:id` - Delete experience

### Images
- `POST /images/upload` - Upload image to S3

### Example: Create Project (with auth)

```bash
# First, get token from auth-service
TOKEN=$(curl -X POST http://localhost:8084/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"password"}' \
  | jq -r '.access_token')

# Create project
curl -X POST http://localhost:8083/api/v1/projects \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My Project",
    "description": "Project description",
    "technologies": ["Go", "Vue.js"]
  }'
```

## Swagger Documentation

When running, Swagger UI is available at:
- `http://localhost:8083/swagger/index.html`

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8083` |
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | Database user | `portfolio_user` |
| `DB_PASSWORD` | Database password | `portfolio_pass` |
| `DB_NAME` | Database name | `portfolio` |
| `AUTH_SERVICE_URL` | Auth service URL | `http://localhost:8084` |
| `S3_ENDPOINT` | MinIO/S3 endpoint | `http://localhost:9000` |
| `S3_ACCESS_KEY` | MinIO access key | `minioadmin` |
| `S3_SECRET_KEY` | MinIO secret key | `minioadmin` |
| `S3_BUCKET` | S3 bucket name | `images` |
| `S3_USE_SSL` | Use SSL for S3 | `false` |

## Development

### Running Tests

```bash
task test
# or
go test ./...
```

### Generating Swagger Docs

```bash
task swagger
# or
swag init -g cmd/api/main.go -o docs
```

### Building

```bash
task build
# or
go build -o bin/admin-api cmd/api/main.go
```

## Authentication

This API validates JWT tokens issued by auth-service. The middleware:
1. Extracts token from `Authorization` header
2. Validates token signature and expiry
3. Injects user information into request context

## Integration

This API is consumed by the admin-web frontend for portfolio content management.

## License

MIT

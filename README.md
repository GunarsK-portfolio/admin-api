# Portfolio Admin API

CRUD REST API for managing portfolio content. Requires authentication via JWT tokens from the auth service.

## Features

- Full CRUD operations for all content
- JWT authentication required
- Validates tokens with auth service
- S3/MinIO integration for images
- Swagger API documentation
- Docker support

## API Endpoints

All endpoints require `Authorization: Bearer <token>` header.

### Profile
- `POST /api/v1/profile` - Update profile

### Work Experience
- `POST /api/v1/experience` - Create work experience
- `PUT /api/v1/experience/:id` - Update work experience
- `DELETE /api/v1/experience/:id` - Delete work experience

### Certifications
- `POST /api/v1/certifications` - Create certification
- `PUT /api/v1/certifications/:id` - Update certification
- `DELETE /api/v1/certifications/:id` - Delete certification

### Miniatures
- `POST /api/v1/miniatures` - Create miniature project
- `PUT /api/v1/miniatures/:id` - Update miniature project
- `DELETE /api/v1/miniatures/:id` - Delete miniature project

### Images
- `DELETE /api/v1/images/:id` - Delete image

### Health
- `GET /api/v1/health` - Service health check (no auth required)

### Documentation
- `GET /swagger/index.html` - Swagger UI

## Quick Start

### Prerequisites

- Go 1.25+
- PostgreSQL 18+
- Auth Service running
- MinIO or S3
- [Task](https://taskfile.dev/installation/) (task runner)
- Docker (optional)

### Local Development (without Docker)

```bash
# Install dependencies
go mod download

# Copy environment file
cp .env.example .env

# Edit .env with your local settings
# Make sure auth-service is running on localhost:8084

# Generate Swagger docs
task swagger

# Run the service
task run

# Or debug in VS Code (F5)
```

### Local Development (with Docker)

```bash
# Start admin API and dependencies
docker-compose up -d

# View logs
docker-compose logs -f admin-api
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Service port | 8083 |
| DB_HOST | PostgreSQL host | localhost |
| DB_PORT | PostgreSQL port | 5432 |
| DB_USER | Database user | portfolio_user |
| DB_PASSWORD | Database password | portfolio_pass |
| DB_NAME | Database name | portfolio |
| S3_ENDPOINT | S3/MinIO endpoint | http://localhost:9000 |
| S3_ACCESS_KEY | S3 access key | minioadmin |
| S3_SECRET_KEY | S3 secret key | minioadmin |
| S3_BUCKET | S3 bucket name | images |
| S3_USE_SSL | Use SSL for S3 | false |
| AUTH_SERVICE_URL | Auth service URL | http://localhost:8084 |

## Authentication

All API endpoints (except `/health`) require a valid JWT token:

```bash
# 1. Login to get token

# Direct access (standalone)
curl -X POST http://localhost:8084/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin123"}'

# Via Traefik (infrastructure setup)
curl -X POST http://localhost:81/auth/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin123"}'

# Response:
# {
#   "access_token": "eyJhbGciOiJI...",
#   "refresh_token": "eyJhbGciOiJI...",
#   "expires_in": 900
# }

# 2. Use token in requests

# Direct access (standalone)
curl -X POST http://localhost:8083/api/v1/profile \
  -H "Authorization: Bearer eyJhbGciOiJI..." \
  -H "Content-Type: application/json" \
  -d '{"full_name": "John Doe", ...}'

# Via Traefik (infrastructure setup)
curl -X POST http://localhost:81/api/v1/profile \
  -H "Authorization: Bearer eyJhbGciOiJI..." \
  -H "Content-Type: application/json" \
  -d '{"full_name": "John Doe", ...}'
```

## API Usage Examples

### Update Profile

```bash
# Direct access (standalone)
curl -X POST http://localhost:8083/api/v1/profile \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "John Doe",
    "title": "Software Engineer",
    "bio": "Passionate developer...",
    "email": "john@example.com"
  }'

# Via Traefik (infrastructure setup)
curl -X POST http://localhost:81/api/v1/profile \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "John Doe",
    "title": "Software Engineer",
    "bio": "Passionate developer...",
    "email": "john@example.com"
  }'
```

### Create Work Experience

```bash
# Direct access (standalone)
curl -X POST http://localhost:8083/api/v1/experience \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{...}'

# Via Traefik (infrastructure setup)
curl -X POST http://localhost:81/api/v1/experience \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "company": "Example Corp",
    "position": "Senior Developer",
    "description": "Led development...",
    "start_date": "2020-01-01",
    "is_current": true,
    "display_order": 1
  }'
```

### Delete Certification

```bash
# Direct access (standalone)
curl -X DELETE http://localhost:8083/api/v1/certifications/1 \
  -H "Authorization: Bearer <token>"

# Via Traefik (infrastructure setup)
curl -X DELETE http://localhost:81/api/v1/certifications/1 \
  -H "Authorization: Bearer <token>"
```

## Development

### Generate Swagger Documentation

```bash
task swagger
git add docs/
git commit -m "Update swagger docs"
```

### Run Tests

```bash
task test
```

### Build Binary

```bash
task build
```

## Docker

### Build Image

```bash
task docker-build
```

### Run with Docker Compose

```bash
task docker-run
```

## Security

- All endpoints protected with JWT authentication
- Tokens validated with auth service on every request
- Tokens expire after 15 minutes
- Use refresh tokens to get new access tokens

## Related Repositories

- [infrastructure](https://github.com/GunarsK-portfolio/infrastructure)
- [database](https://github.com/GunarsK-portfolio/database)
- [auth-service](https://github.com/GunarsK-portfolio/auth-service)
- [admin-web](https://github.com/GunarsK-portfolio/admin-web)

## License

MIT

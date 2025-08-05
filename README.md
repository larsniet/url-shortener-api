# URL Shortener Service

A high-performance URL shortening service built with Go, PostgreSQL, and Docker.

## Features

- 🔗 Shorten long URLs with custom short codes
- 📊 Click tracking and analytics
- 🐳 Docker containerized deployment
- 🚀 High-performance Go backend
- 📦 PostgreSQL database with migrations
- 🔄 Automated CI/CD with GitHub Actions

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for local development)

### Running with Docker

```bash
# Start the services
docker-compose up -d

# The API will be available at http://localhost:8080
```

### Local Development

```bash
# Install dependencies
go mod download

# Start PostgreSQL
docker-compose up -d postgres

# Run migrations
atlas migrate apply --env local --allow-dirty

# Start the server with hot reloading
docker-compose up app
```

## API Documentation

### Create Short URL
```bash
POST /api/v1/urls
Content-Type: application/json

{
  "url": "https://example.com/very-long-url",
  "custom_code": "optional-custom-code"
}
```

### Redirect to Original URL
```bash
GET /{short_code}
# Redirects to original URL
```

### Get URL Analytics
```bash
GET /api/v1/urls/{short_code}/stats
# Returns click count and analytics
```

## Project Structure

```
.
├── cmd/
│   └── server/          # Application entrypoints
├── internal/
│   ├── handler/         # HTTP handlers
│   ├── model/          # Data models
│   ├── services/       # Business logic
│   └── db/             # Database layer & migrations
├── pkg/
│   ├── logger/         # Shared utilities
│   └── utils/          # Helper functions
├── .github/workflows/  # CI/CD pipelines
├── docker-compose.yaml # Development environment
├── Dockerfile          # Development container
└── Dockerfile.prod     # Production container
```

## Database Migrations

We use [Atlas](https://atlasgo.io/) for database migrations:

```bash
# Apply migrations
atlas migrate apply --env local --allow-dirty

# Create new migrations (after updating schema.hcl)
atlas migrate diff migration_name --env local

# Check migration status
atlas migrate status --env local --allow-dirty
```

## Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Deployment

### Docker Production

```bash
# Build production image
docker build -f Dockerfile.prod -t url-shortener .

# Run with environment variables
docker run -p 8080:8080 \
  -e DATABASE_URL=postgres://... \
  url-shortener
```

### GitHub Container Registry

Images are automatically built and published to GitHub Container Registry on:
- Every push to `main` branch
- Git tags (semantic versioning)
- Pull requests (for testing)

```bash
# Pull the latest image
docker pull ghcr.io/OWNER/url-shortener:main

# Pull a specific version
docker pull ghcr.io/OWNER/url-shortener:v1.0.0
```

### CI/CD Pipeline

The GitHub Actions workflow automatically:
1. ✅ Runs tests
2. 🏗️ Builds Docker image with multi-stage build
3. 🐳 Pushes to GitHub Container Registry (GHCR)
4. 🔒 Creates build provenance attestation
5. 🌐 Supports multi-platform builds (amd64, arm64)

## Environment Variables

| Variable       | Description                              | Default  |
| -------------- | ---------------------------------------- | -------- |
| `PORT`         | Server port                              | `8080`   |
| `DATABASE_URL` | PostgreSQL connection string             | Required |
| `LOG_LEVEL`    | Logging level (debug, info, warn, error) | `info`   |

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
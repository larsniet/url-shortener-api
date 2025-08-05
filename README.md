# URL Shortener Service

A high-performance URL shortening service built with Go, PostgreSQL, and Docker.

## Features

- 🔗 Shorten long URLs with custom short codes
- 📊 Click tracking and analytics
- 🐳 Docker containerized deployment
- 🚀 High-performance Go backend
- 📦 PostgreSQL database with migrations

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
go run cmd/migrate/main.go

# Start the server
go run cmd/server/main.go
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
│   └── db/             # Database layer
├── pkg/
│   └── logger/         # Shared utilities
├── docker-compose.yaml # Development environment
└── Dockerfile          # Production container
```

## Development

### Database Migrations

We use [Atlas](https://atlasgo.io/) for database migrations:

```bash
# Create a new migration
atlas migrate diff --env local

# Apply migrations
atlas migrate apply --env local
```

### Testing

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
docker build -t url-shortener .

# Run with environment variables
docker run -p 8080:8080 \
  -e DATABASE_URL=postgres://... \
  url-shortener
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `DATABASE_URL` | PostgreSQL connection string | Required |
| `LOG_LEVEL` | Logging level (debug, info, warn, error) | `info` |

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
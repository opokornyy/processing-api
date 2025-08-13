# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based REST API following clean architecture principles with a layered structure:
- **Controller Layer**: HTTP handlers and routing (`internal/controller/http/v1/`)
- **Use Case Layer**: Business logic (`internal/usecase/`)
- **Repository Layer**: Data access (`internal/repo/`)
- **Configuration**: Application config (`config/`)

The application uses zerolog for structured logging and follows a dependency injection pattern where dependencies flow from main → app → controller → usecase → repository.

## Build and Development Commands

### Container Operations
```bash
# Build container image using Podman
make build-container

# Build and run container (exposes port 8080)
make run
```

### Direct Go Commands
```bash
# Build the application
go build -o main ./cmd/app

# Run the application
go run ./cmd/app

# Download dependencies
go mod download

# Tidy dependencies
go mod tidy
```

## Architecture Details

### Application Flow
1. `cmd/app/main.go` - Entry point, loads config and starts app
2. `internal/app/app.go` - Application bootstrap, wires dependencies and starts HTTP server
3. HTTP requests flow: Router → Handler → UseCase → Repository

### Key Components
- **HTTP Server**: Runs on port 8080 with 10s read/write timeouts
- **Routing**: Uses standard `http.ServeMux` with `/api/v1/` prefix
- **Current Endpoint**: `GET /api/v1/data` returns "Success!" message
- **Dependency Chain**: V1Controller → UseCase → BasicRepo

### TODO Items in Codebase
- Database connection setup (app.go:18)
- Server graceful shutdown (app.go:41)
- Environment-based configuration (config/config.go:10)
- Business logic implementation (usecase/logic.go:20)
- PostgreSQL repository implementation (repo/basic/repo_postgres.go:6)

## Project Structure Notes

- `pkg/` directory exists but is currently empty
- `internal/models/` exists but has no model definitions yet
- Repository layer prepared for PostgreSQL but not implemented
- Configuration struct exists but has no fields defined

The codebase is in early development stage with basic HTTP server functionality and clean architecture scaffolding in place.
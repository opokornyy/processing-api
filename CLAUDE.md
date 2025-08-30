# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based weather processing API following clean architecture principles. The system fetches weather data from external APIs, generates AI-powered recommendations, and processes weather workflow data. The architecture follows domain-driven design with clear separation of concerns:

- **Domain Layer**: Core business entities (`internal/domain/weather/`, `internal/domain/recommendation/`)
- **Use Case Layer**: Business logic orchestration (`internal/usecase/`)
- **Adapter Layer**: External integrations (`internal/adapters/`)
- **Application Layer**: Bootstrap and dependency wiring (`internal/app/`)

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

### Domain Entities
- **WeatherData**: Core weather entity with location, temperature, humidity, and timestamp
- **Recommendation**: AI-generated recommendations linked to weather data
- Both entities use typed IDs (WeatherID, RecommendationID) for type safety

### Use Cases (Business Logic)
- **FetchWeather**: Retrieves weather data from external APIs and persists it
- **GenerateRecommendations**: Creates AI-powered recommendations based on weather
- **ProcessWeatherWorkflow**: Orchestrates the complete weather processing pipeline

### Adapter Interfaces
- **WeatherAPI**: External weather service integration (OpenWeather, RPiWeather)
- **AIService**: AI recommendation generation (OpenAI integration)
- **Repository**: Data persistence with mock and PostgreSQL implementations
- **Notification**: Email notification service

### Application Flow
1. `cmd/app/main.go` - Entry point, loads config and starts application
2. `internal/app/app.go` - Application bootstrap (currently minimal, needs HTTP server setup)
3. Configuration uses zerolog for structured logging with Unix timestamp format

### Current Implementation Status
- Domain entities are fully implemented with proper constructors
- Use case interfaces (ports) are defined but implementation is placeholder
- Adapter interfaces are defined but implementations are stubs
- HTTP handlers exist but are not wired up yet
- Application bootstrap is minimal and needs HTTP server setup

### Key Dependencies
- `github.com/rs/zerolog` - Structured logging library
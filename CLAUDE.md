# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a full-stack application template using Buf for API schema management, Go with Connect RPC for the backend, and React with TypeScript for the frontend. The project demonstrates a complete todo application with proper API codegen, database integration, and monitoring.

## Architecture

- **Schema-first development**: Protocol Buffers definitions in `_schema/protos/` generate both Go and TypeScript SDKs
- **Backend**: Go server using Connect RPC framework, PostgreSQL database with GORM for ORM
- **Frontend**: React + TypeScript app consuming the generated web SDK
- **Monitoring**: Prometheus + Grafana stack for observability
- **Workspace structure**: NPM workspaces for frontend + web SDK, Go workspace for backend + Go SDK

## Development Commands

### Environment Setup
```bash
# Install dependencies
nvm install
npm ci

# Start monitoring stack (Grafana at :3000, Prometheus at :9090)
make monitoring-up

# Start database and services
make stack-up
```

### Code Generation
```bash
# Generate API SDKs from protobuf schemas
make api-codegen
```

### Development Servers
```bash
# Start backend server
make run-backend-server

# Start frontend development server
npm run frontend:dev

# Start backend with hot reload (modd)
make backend-run

# Start backend in debug mode
make backend-debug
```

### Build and Quality
```bash
# Build all workspaces
npm run build

# Lint all workspaces
npm run lint

# Format frontend code
npm run format

# Lint Go backend with golangci-lint (Docker-based)
make lint

# Auto-fix Go linting issues
make lint-fix
```

## Key Files and Directories

- `_schema/`: Protocol buffer definitions and buf configuration
- `api/`: Generated SDKs (go-sdk and web-sdk)
- `backend/`: Go server with Connect RPC handlers and database layer
- `frontend/`: React application
- `_harness/`: Docker compose files for local development services
- `makefiles/`: Modular Makefile organization

## Database Integration

The project uses PostgreSQL with GORM for ORM. Database connection is configured via environment variables in `backend/internal/config/config.go`. The database service runs in Docker via `make stack-up`.

## API Development

1. Define services in `_schema/protos/`
2. Run `make api-codegen` to generate SDKs
3. Implement handlers in `backend/internal/api-*/`
4. Use generated types in frontend via web SDK
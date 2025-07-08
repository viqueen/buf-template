# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a full-stack application template using Buf for API schema management, Go with Connect RPC for the backend, and React with TypeScript for the frontend. The project demonstrates a complete todo application with proper API codegen, database integration, and monitoring.

## Architecture

- **Schema-first development**: Protocol Buffers definitions in `_schema/protos/` generate both Go and TypeScript SDKs
- **Backend**: Go server using Connect RPC framework, PostgreSQL database with SQLC for type-safe queries
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

# Run database migrations
make migrations
```

### Code Generation
```bash
# Generate API SDKs from protobuf schemas
make api-codegen

# Generate Go database code from SQL queries
make sqlc-codegen
```

### Development Servers
```bash
# Start backend server
make run-backend-server

# Start frontend development server
npm run frontend:dev
```

### Build and Quality
```bash
# Build all workspaces
npm run build

# Lint all workspaces
npm run lint

# Format frontend code
npm run format
```

## Key Files and Directories

- `_schema/`: Protocol buffer definitions and buf configuration
- `api/`: Generated SDKs (go-sdk and web-sdk)
- `backend/`: Go server with Connect RPC handlers and database layer
- `frontend/`: React application
- `_harness/`: Docker compose files for local development services
- `makefiles/`: Modular Makefile organization

## Database Integration

The project uses PostgreSQL with Flyway migrations (`backend/data/schema/`) and SQLC for type-safe query generation. Database queries are defined in `backend/data/queries/` and generate Go code in `backend/internal/store/gendb/`.

## API Development

1. Define services in `_schema/protos/`
2. Run `make api-codegen` to generate SDKs
3. Implement handlers in `backend/internal/api-*/`
4. Use generated types in frontend via web SDK
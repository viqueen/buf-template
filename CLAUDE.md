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
npm run harness start-monitoring

# Start database and services
npm run harness start-stack
```

### Code Generation
```bash
# Generate API SDKs from protobuf schemas
npm run schema codegen
```

### Development Servers
```bash
# Start backend server with hot reload (modd)
npm run backend start

# Start frontend development server
npm run frontend dev

# Start backend in debug mode
npm run backend debug
```

### Build and Quality
```bash
# Build all workspaces
npm run build

# Lint all workspaces
npm run lint

# Format frontend code
npm run format

# Lint Go backend
npm run backend lint

# Auto-fix Go linting issues
npm run backend lint:fix
```

### Stopping Services
```bash
# Stop monitoring stack
npm run harness stop-monitoring

# Stop database and services
npm run harness stop-stack
```

## Key Files and Directories

- `_schema/`: Protocol buffer definitions and buf configuration
- `api/`: Generated SDKs (go-sdk and web-sdk)
- `backend/`: Go server with Connect RPC handlers and database layer
- `frontend/`: React application
- `_harness/`: Docker compose files for local development services

## Database Integration

The project uses PostgreSQL with GORM for ORM. Database connection is configured via environment variables in `backend/internal/config/config.go`. The database service runs in Docker via `npm run harness start-stack`.

## API Development Workflow

1. Define services in `_schema/protos/`
2. Run `npm run schema codegen` to generate SDKs
3. Implement handlers in `backend/internal/api-*/`
4. Use generated types in frontend via web SDK imports from `@labset/buf-template-web-sdk`

## Connect RPC Handler Pattern

Backend handlers follow this pattern:
```go
func (s *Service) MethodName(
    ctx context.Context,
    req *connect.Request[protov1.MethodRequest],
) (*connect.Response[protov1.MethodResponse], error) {
    // Implementation
}
```

## Frontend API Client Usage

Import and use generated clients:
```typescript
import { createConnectTransport } from "@connectrpc/connect-web";
import { createPromiseClient } from "@connectrpc/connect";
import { TodoService } from "@labset/buf-template-web-sdk/todo/v1/todo_connect";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

const client = createPromiseClient(TodoService, transport);
```
# buf-template

This is a template for using the [buf](https://buf.build) build system to define APIs and
implement them with Go using Connect rpc framework.

## development

### setup

- install dependencies

```bash
nvm install
npm ci
```

- run api codegen

```bash
make api-codegen
```

- invite claude

```bash
npm run claude
```

### environment

```bash
make monitoring-up
```

- open grafana at http://localhost:3000 and login with `admin`/`admin`
- open prometheus at http://localhost:9090

```bash
make stack-up
```

### backend

- run backend server

```bash
make run-backend-server
```

### frontend

- run frontend server

```bash
npm run frontend:dev
```
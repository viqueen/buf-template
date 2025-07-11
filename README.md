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
npm run schema codegen
```

- invite claude

```bash
npm run claude
```

### environment

```bash
npm run harness start-monitoring
```

- open grafana at http://localhost:3000 and login with `admin`/`admin`
- open prometheus at http://localhost:9090

```bash
npm run harness start-stack
```

### backend

- run backend server

```bash
npm run backend start
```

### frontend

- run frontend server

```bash
npm run frontend dev
```
# buf-template

This is a template for using the [buf](https://buf.build) build system to define APIs and
implement them with Go using Connect rpc framework.

## setup

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

- run sql migrations with flyway

```bash
make migrations
```

- run api codegen

```bash
make codegen
```

- run sqlc codegen

```bash
make sqlc
```

- run backend server

```bash
make run-backend-server
```
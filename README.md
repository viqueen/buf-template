## buf-template

- This is a template for using the [buf](https://buf.build) build system. 
- It includes codegen for Go with Connect rpc.

### development setup

- start harness

```bash
make harness-up
```

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
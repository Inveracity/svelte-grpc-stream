# Svelte + gRPC

:warning: work in progress :warning:

Learning some svelte and trying to stream realtime updates to it from a gRPC server.

## Requirements

- [buf](buf.build) to generate code from protobuffers
- [node](nodejs.org)
- [golang](go.dev)

```sh
go mod tidy
npm install
buf generate
```

```sh
docker compose up --build
```

```bash
npm run dev -- --open
```

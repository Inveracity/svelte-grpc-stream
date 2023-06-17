# Svelte + gRPC

:warning: work in progress :warning:

Learning some svelte and trying to stream realtime updates to it from a gRPC server.

## Requirements

- [node](nodejs.org) to run svelte app
- [golang](go.dev) to run golang server
- [buf](buf.build) to generate code from protobuffers

## Dev

```sh
go mod tidy
npm install
make proto
```

```sh
docker compose up --build
```

```bash
npm run dev -- --open
```

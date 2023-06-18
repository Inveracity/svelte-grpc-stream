# Svelte + gRPC stream

:warning: work in progress :warning:

Learning some svelte and trying to stream realtime updates to it from a gRPC server.

## Infrastructure

This project consists of 5 infrastructure components
- Traefik: ingress controller and gRPC-Web proxy
- NATS: for queueing notifications that need to be relayed
- API (python): For putting notifications on the queue
- Relay (golang): Picks up messages from the queue and forwards them to a gRPC stream to the frontend
- Frontend (svelte & typecript): For sending notifications via the API and recieving notifications via the Relay.

## Requirements

- [node](nodejs.org) to run svelte app
- [golang](go.dev) to run the golang server
- [buf](buf.build) to generate code from protobuffers
- [python 3.11](python.org) to run the python API

## Dev

TODO: write a more comprehensive guide to getting set up

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

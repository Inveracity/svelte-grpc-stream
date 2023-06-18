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

- [node 20.3.0](nodejs.org) to run svelte app
- [golang 1.20](go.dev) to run the golang server
- [python 3.11.4](python.org) to run the python API
- [buf 1.21.0](buf.build) to generate code from protobuffers
- [docker 24.0.2](docker.com) to run services in docker compose

## Dev

- [relay](relay/readme.md)
- [api](api/readme.md)
- [frontend](frontend/readme.md)

## Ingress

In your hosts file set:

```plaintext
127.0.0.1 frontend.docker.localhost
127.0.0.1 relay.docker.localhost
127.0.0.1 api.docker.localhost
```

After running `docker compose up`, <http://frontend.docker.localhost> should now be available via Traefik.
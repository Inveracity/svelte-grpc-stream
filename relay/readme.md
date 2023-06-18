# Relay

This doc describes how to develop and run the relay service

## Install

```sh
# Install dependencies
cd relay
go mod tidy

# Generate gRPC code
cd ..
make proto

# Run the relay server
cd relay
go run cmd/relay/main.go
```

The relay server will not connect to NATS until a user requests a queue to be created.

from the root of the project run NATS via docker compose.

```sh
docker compose up -d nats
```

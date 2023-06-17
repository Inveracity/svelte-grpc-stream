package main

import (
	"flag"

	"github.com/inveracity/svelte-grpc-stream/internal/server"
)

var (
	port = flag.Int("port", 50051, "The server port")
	nats = flag.String("nats", "nats:4222", "The nats server")
)

func main() {
	flag.Parse()
	server.Run(*port, *nats)
}

package main

import (
	"flag"

	"github.com/inveracity/svelte-grpc-stream/internal/relay"
)

var (
	port  = flag.Int("port", 50051, "The server port")
	nats  = flag.String("nats", "nats:4222", "The nats server")
	redis = flag.String("redis", "redis:6379", "The redis server")
)

func main() {
	flag.Parse()
	relay := relay.NewRelay(*port, *nats, *redis)
	err := relay.Run()
	if err != nil {
		panic(err)
	}
}

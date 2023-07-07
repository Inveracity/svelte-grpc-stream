package main

import (
	"flag"
	"os"

	"github.com/inveracity/svelte-grpc-stream/internal/relay"
)

var (
	port    = flag.Int("port", 50051, "The server port")
	nats    = flag.String("nats", "nats:4222", "The nats server")
	redis   = flag.String("redis", "redis:6379", "The redis server")
	pbURL   = flag.String("pocketbase", "http://pocketbase:8090", "The pocketbase server")
	pbAdmin = os.Getenv("PB_USER")
	pbPass  = os.Getenv("PB_PASS")
)

func main() {
	flag.Parse()
	relay := relay.NewRelay(*port, *nats, *redis, *pbURL, pbAdmin, pbPass)
	err := relay.Run()
	if err != nil {
		panic(err)
	}
}

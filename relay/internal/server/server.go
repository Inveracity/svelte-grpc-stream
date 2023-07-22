package server

import (
	"github.com/inveracity/svelte-grpc-stream/internal/cache"
	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/chat/v1"
	"github.com/inveracity/svelte-grpc-stream/internal/queue"
)

var ServerID = "myserver" // TODO: ServerID should be a configurable value

type Server struct {
	pb.UnimplementedChatServiceServer
	cache    *cache.Cache
	queue    *queue.Queue
	streamid string
	natsURL  string
}

func NewServer(natsURL string, cache *cache.Cache) *Server {
	return &Server{
		cache:   cache,
		natsURL: natsURL,
	}
}

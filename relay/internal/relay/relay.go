package relay

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/inveracity/svelte-grpc-stream/internal/cache"
	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/notifications/v1"
	"github.com/inveracity/svelte-grpc-stream/internal/queue"
	"github.com/inveracity/svelte-grpc-stream/internal/server"

	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Relay struct {
	ctx    context.Context
	server *server.Server
	nats   *nats.Conn
	port   int
}

func NewRelay(ctx context.Context, port int, natsURL string, redisURL string) *Relay {
	natsConn, err := nats.Connect(natsURL)
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})

	cache := cache.NewCache(ctx, redisClient)

	messages := make(chan nats.Msg, 64)
	queue := queue.NewQueue(ctx, natsConn, &messages)

	grpcServer := server.NewServer(ctx, cache, queue)
	return &Relay{
		port:   port,
		server: grpcServer,
		nats:   natsConn,
		ctx:    ctx,
	}
}

func (r *Relay) Run() error {
	lis, err := net.Listen("tcp4", fmt.Sprintf(":%d", r.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, r.server)

	log.Printf("GRPC: server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

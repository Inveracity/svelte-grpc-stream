package relay

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/inveracity/svelte-grpc-stream/internal/cache"
	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/chat/v1"
	"github.com/inveracity/svelte-grpc-stream/internal/server"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Relay struct {
	server *server.Server
	port   int
}

func NewRelay(port int, natsURL string, redisURL string) *Relay {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})

	cache := cache.NewCache(redisClient)

	grpcServer := server.NewServer(cache)
	return &Relay{
		port:   port,
		server: grpcServer,
	}
}

func (r *Relay) Run() error {
	lis, err := net.Listen("tcp4", fmt.Sprintf(":%d", r.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             time.Duration(1 * time.Second),
			PermitWithoutStream: true, // Allow pings even when there are no active streams
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    time.Duration(2 * time.Hour),
			Timeout: time.Duration(20 * time.Second),
		}),
	)
	pb.RegisterChatServiceServer(s, r.server)

	log.Printf("GRPC: server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

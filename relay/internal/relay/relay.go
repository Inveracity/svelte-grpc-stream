package relay

import (
	"fmt"
	"log"
	"net"

	"github.com/inveracity/svelte-grpc-stream/internal/auth"
	"github.com/inveracity/svelte-grpc-stream/internal/cache"
	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/chat/v1"
	"github.com/inveracity/svelte-grpc-stream/internal/server"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Relay struct {
	server  *server.Server
	port    int
	pbURL   string
	pbAdmin string
	pbPass  string
}

func NewRelay(port int, natsURL, redisURL, pbURL, pbAdmin, pbPass string) *Relay {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})

	cache := cache.NewCache(redisClient)

	grpcServer := server.NewServer(natsURL, cache)
	return &Relay{
		port:    port,
		server:  grpcServer,
		pbURL:   pbURL,
		pbAdmin: pbAdmin,
		pbPass:  pbPass,
	}
}

func (r *Relay) Run() error {
	lis, err := net.Listen("tcp4", fmt.Sprintf(":%d", r.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authMgr := auth.New(r.pbURL, r.pbAdmin, r.pbPass)
	interceptor := server.NewAuthInterceptor(authMgr)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	pb.RegisterChatServiceServer(s, r.server)

	log.Printf("GRPC: server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

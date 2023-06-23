package server

import (
	"log"
	"sync"

	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/notifications/v1"
	"github.com/nats-io/nats.go"
)

func (s *server) Subscribe(in *pb.SubscribeRequest, srv pb.NotificationService_SubscribeServer) error {
	ctx := srv.Context()
	var wg sync.WaitGroup

	log.Printf("GRPC: user %s connected to channel %s", in.ChannelId, in.UserId)

	// Create a channel to pass messages from NATS to this gRPC server stream
	eventChannel := make(chan nats.Msg)

	// Pass the go channel into the NATS loop
	go NatsSub(ctx, s.nats, in.ChannelId, &eventChannel)

	// send a "connected" message to the client to tell the client it successfully connected
	VerifySubscription(srv, in)

	for {
		select {
		case <-ctx.Done():
			log.Printf("disconnected %s", in.ChannelId)
			return nil
		default:
			for event := range eventChannel {
				wg.Add(1)
				go func(event nats.Msg) {
					defer wg.Done()
					forwardEventToClient(event, srv)
				}(event)
			}
		}
		wg.Wait()
	}
}

func VerifySubscription(srv pb.NotificationService_SubscribeServer, in *pb.SubscribeRequest) {
	srv.Send(&pb.SubscribeResponse{ChannelId: in.ChannelId, UserId: "server", Text: "connected"})
}

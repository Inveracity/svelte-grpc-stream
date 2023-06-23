package server

import (
	"log"
	"sync"

	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/notifications/v1"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/encoding/protojson"
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
	verifySubscription(srv, in)

	// Receive messages from the NATS loop and forward them to the client
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
					relay(event, srv)
				}(event)
			}
		}
		wg.Wait()
	}
}

func verifySubscription(srv pb.NotificationService_SubscribeServer, in *pb.SubscribeRequest) {
	srv.Send(&pb.SubscribeResponse{ChannelId: in.ChannelId, UserId: "server", Text: "connected"})
}

// Send messages from NATS to the gRPC client
func relay(
	event nats.Msg,
	srv pb.NotificationService_SubscribeServer,
) {
	var notification pb.SubscribeResponse

	log.Printf("forwarding event from nats to grpc: %s", string(event.Data))
	// unmarshal the nats message into a protobuf message
	j := protojson.UnmarshalOptions{}
	if err := j.Unmarshal(event.Data, &notification); err != nil {
		log.Printf("unmarshal error %v", err)
		return
	}

	if err := srv.Send(&notification); err != nil {
		log.Printf("send error %v", err)
		return
	}

	// Ack the NATS message so it's not sent again
	event.Ack()
}

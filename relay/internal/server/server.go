package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/notifications/v1"
	"github.com/nats-io/nats.go"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
	nats string
}

func Run(port int, nats string) {
	lis, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, &server{nats: nats})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Send messages from NATS to the gRPC client
func forwardEventToClient(
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

// Subscribe to NATS
// TODO: make this a separate package
func NatsSub(ctx context.Context, url, channelId string, events *chan nats.Msg) error {
	log.Printf("NATS: subscribing to channel: %s", channelId)
	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}

	subject := "events." + channelId
	msgChan := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe(subject, msgChan)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Printf("NATS: unsubscribing disconnected client: %s", channelId)
			sub.Unsubscribe()
			return nil

		default:
			msg, err := sub.NextMsgWithContext(ctx)
			if err != nil {
				log.Printf("next message error: %v", err)
				continue
			}

			// the nats message is sent back to the gRPC handler via the events channel, and will be "Ack()"ed there
			*events <- *msg
		}
	}
}

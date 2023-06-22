package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/notifications/v1"
	"github.com/nats-io/nats.go"

	"google.golang.org/grpc"
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

func (s *server) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	log.Printf("user: %s sent: %s to channel: %s", in.Notification.UserId, *in.Notification.Text, in.Notification.ChannelId)
	nc, err := nats.Connect(s.nats)
	if err != nil {
		return nil, err
	}

	subject := "events." + in.Notification.ChannelId
	msg := nats.NewMsg(subject)
	msg.Data = []byte(fmt.Sprintf(
		`{"ChannelId":"%s", "userId": "%s", "text":"%s"}`,
		in.Notification.ChannelId,
		in.Notification.UserId,
		*in.Notification.Text,
	))

	if err := nc.PublishMsg(msg); err != nil {
		return nil, err
	}

	return &pb.SendResponse{}, nil
}

func (s *server) Subscribe(in *pb.SubscribeRequest, srv pb.NotificationService_SubscribeServer) error {
	ctx := srv.Context()
	var wg sync.WaitGroup

	log.Printf("connected: %s", in.Notification.ChannelId)

	eventChannel := make(chan nats.Msg)

	go NatsSub(ctx, s.nats, in.Notification.ChannelId, &eventChannel)

	// send a "connected" message to the client
	queueName := fmt.Sprintf("events.%s", in.Notification.ChannelId)
	connectResponse := *nats.NewMsg(queueName)
	connectResponse.Data = []byte(`{"ChannelId":"` + in.Notification.ChannelId + `", "userId": "server", "text":"connected"}`)
	forwardEventToClient(connectResponse, srv)

	for {
		select {
		case <-ctx.Done():
			log.Printf("disconnected %s", in.Notification.ChannelId)
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

// Send messages from NATS to the gRPC client
func forwardEventToClient(
	event nats.Msg,
	srv pb.NotificationService_SubscribeServer,
) {
	var notification pb.Notification

	// unmarshal the nats message into a protobuf message
	if err := json.Unmarshal(event.Data, &notification); err != nil {
		log.Printf("unmarshal error %v", err)
		return
	}

	resp := &pb.SubscribeResponse{Notification: &notification}
	if err := srv.Send(resp); err != nil {
		log.Printf("send error %v", err)
		return
	}

	// Ack the NATS message so it's not sent again
	event.Ack()
}

// Subscribe to NATS
// TODO: make this a separate package
func NatsSub(ctx context.Context, url, channelId string, events *chan nats.Msg) error {
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
			log.Printf("unsubscribing disconnected client: %s", channelId)
			sub.Unsubscribe()
			return nil

		default:
			msgs, err := sub.Fetch(1, nats.MaxWait(1*time.Second))
			if err != nil {
				continue
			}
			if len(msgs) == 0 {
				continue
			}
			msg := msgs[0]
			// the nats message is sent back to the gRPC handler via the events channel, and will be "Ack()"ed there
			*events <- *msg
		}
	}
}

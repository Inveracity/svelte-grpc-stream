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

func (s *server) Notify(in *pb.SubscribeRequest, srv pb.NotificationService_NotifyServer) error {
	ctx := srv.Context()
	var wg sync.WaitGroup

	log.Printf("connected: %s", in.Subid)

	eventChannel := make(chan nats.Msg)

	go subscribe(ctx, s.nats, in.Subid, &eventChannel)

	for {
		select {
		case <-ctx.Done():
			log.Printf("disconnected %s", in.Subid)
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

func forwardEventToClient(
	event nats.Msg,
	srv pb.NotificationService_NotifyServer,
) {

	var data pb.Notification
	if err := json.Unmarshal(event.Data, &data); err != nil {
		log.Printf("unmarshal error %v", err)
		return
	}

	n := pb.Notification{
		Subid:  data.Recipient,
		Text:   data.Text,
		Sender: data.Sender,
	}

	resp := &pb.NotificationServiceNotifyResponse{Notifications: &n}
	if err := srv.Send(resp); err != nil {
		log.Printf("send error %v", err)
		return
	}
	event.Ack()
}

func subscribe(ctx context.Context, url, subscriberId string, events *chan nats.Msg) error {
	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}

	js, err := nc.JetStream()

	if err != nil {
		return err
	}

	cfg := &nats.StreamConfig{
		Name:      "EVENTS",
		Retention: nats.WorkQueuePolicy,
		Subjects:  []string{"events.>"},
	}

	js.AddStream(cfg)

	subject := "events." + subscriberId
	sub, err := js.PullSubscribe(subject, subscriberId, nats.BindStream(cfg.Name))
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Printf("unsubscribing disconnected client: %s", subscriberId)
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

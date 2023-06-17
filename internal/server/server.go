package server

import (
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/notifications/v1"
	"github.com/inveracity/svelte-grpc-stream/internal/queue"
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

func (s *server) Notify(in *pb.NotifyRequest, srv pb.NotificationService_NotifyServer) error {
	ctx := srv.Context()
	log.Printf("opening stream for subscriberId: %s", in.Id)

	eventChannel := make(chan nats.Msg)

	go queue.Client(ctx, s.nats, in.Id, &eventChannel)

	var wg sync.WaitGroup

	for {
		select {
		case <-ctx.Done():
			log.Printf("client %s disconnected", in.Id)
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

func forwardEventToClient(event nats.Msg, srv pb.NotificationService_NotifyServer) {
	n := pb.Notification{Id: "1", Msg: string(event.Data)}
	resp := &pb.NotificationServiceNotifyResponse{Notifications: &n}
	if err := srv.Send(resp); err != nil {
		log.Printf("send error %v", err)
		return
	}
	event.Ack()
}

package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/inveracity/svelte-grpc-stream/internal/cache"
	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/notifications/v1"
	"github.com/inveracity/svelte-grpc-stream/internal/queue"
	"github.com/nats-io/nats.go"

	"google.golang.org/protobuf/encoding/protojson"
)

type Server struct {
	pb.UnimplementedNotificationServiceServer
	cache *cache.Cache
	queue *queue.Queue
	ctx   context.Context
}

func NewServer(ctx context.Context, cache *cache.Cache, queue *queue.Queue) *Server {
	return &Server{
		ctx:   ctx,
		cache: cache,
		queue: queue,
	}
}

func (s *Server) Subscribe(in *pb.SubscribeRequest, srv pb.NotificationService_SubscribeServer) error {
	var wg sync.WaitGroup

	log.Printf("GRPC: user %s connected to channel %s", in.UserId, in.ChannelId)

	// Pass the go channel into the NATS loop
	go s.queue.Subscribe(in.ChannelId)

	// send a "connected" message to the client to tell the client it successfully connected
	verifySubscription(srv, in)

	pastMessage, err := s.cache.GetFrom(in.ChannelId, in.LastTs, "+inf")
	if err != nil {
		return err
	}

	for _, message := range pastMessage {
		var notification pb.Notification
		j := protojson.UnmarshalOptions{}
		if err := j.Unmarshal([]byte(message), &notification); err != nil {
			log.Printf("unmarshal error %v", err)
			return err
		}
		if err := srv.Send(&notification); err != nil {
			log.Printf("send error %v", err)
			return err
		}
	}

	// Receive messages from the NATS loop and forward them to the client
	for {
		select {
		case <-s.ctx.Done():
			log.Printf("disconnected %s", in.ChannelId)
			return nil
		default:
			for message := range *s.queue.Messages {
				wg.Add(1)
				go func(message nats.Msg) {
					defer wg.Done()
					relay(message, srv)
				}(message)
			}
		}
		wg.Wait()
	}
}

// Send receives a message from the client and publishes it to the NATS server
func (s *Server) Send(ctx context.Context, in *pb.Notification) (*pb.SendResponse, error) {
	log.Printf("GRPC: user: %s sent: %s to channel: %s", in.UserId, in.Text, in.ChannelId)

	// Override timstamp
	in.Ts = fmt.Sprint(time.Now().UnixNano())

	msg := nats.NewMsg(in.ChannelId)

	payload, err := ProtoToJSON(in)
	if err != nil {
		return nil, err
	}

	msg.Data = payload

	if err := s.cache.Set(in.ChannelId, string(payload)); err != nil {
		return nil, err
	}

	if err := s.queue.Publish(in.ChannelId, payload); err != nil {
		return nil, err
	}

	return &pb.SendResponse{Ok: true, Error: ""}, nil
}

func verifySubscription(srv pb.NotificationService_SubscribeServer, in *pb.SubscribeRequest) {
	srv.Send(&pb.Notification{
		ChannelId: in.ChannelId,
		UserId:    "server",
		Text:      "connected",
		Ts:        "0",
	})
}

// Send messages from NATS to the gRPC client
func relay(message nats.Msg, srv pb.NotificationService_SubscribeServer) {

	// Convert JSON message to Notification object
	notification, err := JSONToProto(message.Data)
	if err != nil {
		log.Printf("unmarshal error %v", err)
		return
	}

	// Override the timestamp with the current time
	notification.Ts = fmt.Sprint(time.Now().UnixNano())
	log.Printf("NATS->GRPC: %s", string(notification.Ts))

	if err := srv.Send(notification); err != nil {
		log.Printf("send error %v", err)
		return
	}

	// Ack the NATS message so it's not sent again
	message.Ack()
}

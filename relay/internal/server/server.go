package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/notifications/v1"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
	redis *redis.Client
	nats  string
}

func Run(port int, nats string, redisAddr string) {
	lis, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	server := &server{
		redis: redisClient,
		nats:  nats,
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, server)
	log.Printf("GRPC: server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Subscribe(in *pb.SubscribeRequest, srv pb.NotificationService_SubscribeServer) error {
	ctx := srv.Context()
	var wg sync.WaitGroup

	log.Printf("GRPC: user %s connected to channel %s", in.UserId, in.ChannelId)

	// Create a channel to pass messages from NATS to this gRPC server stream
	eventChannel := make(chan nats.Msg)

	// Pass the go channel into the NATS loop
	go NatsSub(ctx, s.nats, in.ChannelId, &eventChannel)

	// send a "connected" message to the client to tell the client it successfully connected
	verifySubscription(srv, in)
	pastMessage, err := s.LoadHistory(ctx, in.ChannelId, in.LastTs)
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

// Send receives a request from the client and publishes it to the NATS server
func (s *server) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	log.Printf("GRPC: user: %s sent: %s to channel: %s", in.UserId, in.Text, in.ChannelId)
	nc, err := nats.Connect(s.nats)
	if err != nil {
		return nil, err
	}

	subject := "events." + in.ChannelId
	msg := nats.NewMsg(subject)

	j := protojson.MarshalOptions{UseProtoNames: true}
	payload, err := j.Marshal(in)
	if err != nil {
		return nil, err
	}

	msg.Data = payload

	s.Store(ctx, string(payload), in.ChannelId)
	if err := nc.PublishMsg(msg); err != nil {
		return nil, err
	}

	return &pb.SendResponse{Ok: true, Error: ""}, nil
}

func verifySubscription(srv pb.NotificationService_SubscribeServer, in *pb.SubscribeRequest) {
	srv.Send(&pb.Notification{ChannelId: in.ChannelId, UserId: "server", Text: "connected"})
}

// Send messages from NATS to the gRPC client
func relay(event nats.Msg, srv pb.NotificationService_SubscribeServer) {
	var notification pb.Notification

	log.Printf("forwarding event from nats to grpc: %s", string(event.Data))
	// unmarshal the nats message into a protobuf message
	j := protojson.UnmarshalOptions{}
	if err := j.Unmarshal(event.Data, &notification); err != nil {
		log.Printf("unmarshal error %v", err)
		return
	}

	notification.Ts = fmt.Sprint(time.Now().UnixNano())

	if err := srv.Send(&notification); err != nil {
		log.Printf("send error %v", err)
		return
	}

	// Ack the NATS message so it's not sent again
	event.Ack()
}

// Store messages in Redis
func (s *server) Store(ctx context.Context, payload, channelid string) error {
	z := redis.ZAddArgs{
		Members: []redis.Z{
			{
				Score:  float64(time.Now().UnixNano()),
				Member: payload,
			},
		},
	}

	ret := s.redis.ZAddArgs(ctx, "events:"+channelid, z)

	if ret.Err() != nil {
		return ret.Err()
	}

	return nil
}

// LoadHistory expects a client to know the timestamp of the last message it received in order to retreive all unread messages.
// If the client has never received a message, it should pass 0 as the LastTimestamp.
func (s *server) LoadHistory(ctx context.Context, channelid, LastTimestamp string) ([]string, error) {
	// Construct a query to get all messages since LastTimestamp
	q := redis.ZRangeArgs{
		Key:     "events:" + channelid,
		Start:   LastTimestamp,
		Stop:    "+inf",
		ByScore: true,
		Count:   100,
	}

	messages, err := s.redis.ZRangeArgs(ctx, q).Result()
	if err != nil {
		return []string{}, err
	}

	return messages, nil
}

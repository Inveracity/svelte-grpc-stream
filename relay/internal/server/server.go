package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/inveracity/svelte-grpc-stream/internal/cache"
	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/chat/v1"
	"github.com/inveracity/svelte-grpc-stream/internal/queue"

	"google.golang.org/protobuf/encoding/protojson"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	cache *cache.Cache
	queue *queue.Queue
}

func NewServer(cache *cache.Cache, queue *queue.Queue) *Server {
	return &Server{
		cache: cache,
		queue: queue,
	}
}

func (s *Server) Connect(in *pb.ConnectRequest, srv pb.ChatService_ConnectServer) error {
	log.Printf("GRPC: user %s connected to server %s", in.UserId, in.ServerId)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Pass the go channel into the NATS loop
	wg.Add(1)
	go s.queue.Subscribe(ctx, in.ServerId)

	// send a "connected" message to the client to tell the client it successfully connected
	s.verifyConnection(srv, in)

	// getPastMessages
	if err := s.getPastMessages(srv, in); err != nil {
		log.Printf("error getting past messages: %v", err)
		return err
	}

	// Receive messages from the NATS loop and forward them to the client
	for {
		select {
		case <-srv.Context().Done():
			log.Printf("GRPC: %s disconnected from %s. Server context cancelled.", in.UserId, in.ServerId)
			cancel()
			return nil
		case <-ctx.Done():
			log.Print("GRPC: Client loop was closed")
		default:
			wg.Add(1)
			go Forward(ctx, in, srv, s, cancel)
			wg.Wait()
		}
	}
}

// Send receives a message from the client and publishes it to the NATS server
func (s *Server) Send(ctx context.Context, in *pb.ChatMessage) (*pb.SendResponse, error) {
	log.Printf("GRPC: user: %s sent: %s to channel: %s on server: myserver", in.UserId, in.Text, in.ChannelId)

	// Override timstamp
	in.Ts = fmt.Sprint(time.Now().UnixNano())

	msg := nats.NewMsg("myserver")

	payload, err := ProtoToJSON(in)
	if err != nil {
		return nil, err
	}

	msg.Data = payload

	if err := s.cache.Set("myserver", string(payload)); err != nil {
		log.Printf("error writing message to cache: %v", err)
		return nil, err
	}

	if err := s.queue.Publish("myserver", payload); err != nil {
		log.Printf("error publishing message to queue: %v", err)
		return nil, err
	}

	return &pb.SendResponse{Ok: true, Error: ""}, nil
}

func (s *Server) verifyConnection(srv pb.ChatService_ConnectServer, in *pb.ConnectRequest) {
	srv.Send(&pb.ChatMessage{
		ChannelId: "system", // system information channel
		UserId:    "server",
		Text:      "connected",
		Ts:        "0",
	})
}

// Send messages from NATS to the gRPC client
func relay(message nats.Msg, srv pb.ChatService_ConnectServer, cancel context.CancelFunc) {

	// Convert JSON message to Notification object
	chatMsg, err := JSONToProto(message.Data)
	if err != nil {
		log.Printf("unmarshal error %v", err)
		return
	}

	// Override the timestamp with the current time
	chatMsg.Ts = fmt.Sprint(time.Now().UnixNano())
	log.Printf("NATS->GRPC:%s %s %s: %s", chatMsg.ChannelId, chatMsg.UserId, chatMsg.Ts, chatMsg.Text)

	if err := srv.Send(chatMsg); err != nil {
		log.Printf("an error occurred while relaying new chat messages: %v", err)
		cancel()
		return
	}
}

func (s *Server) getPastMessages(srv pb.ChatService_ConnectServer, in *pb.ConnectRequest) error {
	pastMessages, err := s.cache.GetFrom(in.ServerId, in.LastTs, "+inf")
	if err != nil {
		return err
	}

	for _, message := range pastMessages {
		var chatmsg pb.ChatMessage
		j := protojson.UnmarshalOptions{}

		// Convert JSON message to Notification object
		if err := j.Unmarshal([]byte(message), &chatmsg); err != nil {
			log.Printf("unmarshal error %v", err)
			return err
		}

		// Send the message to the client
		if err := srv.Send(&chatmsg); err != nil {
			log.Printf("an error occurred sending history to client: %v", err)
			close(*s.queue.Messages)
			return err
		}
	}

	return nil
}

func Forward(ctx context.Context, in *pb.ConnectRequest, srv pb.ChatService_ConnectServer, s *Server, cancel context.CancelFunc) error {
	for {
		select {
		case <-ctx.Done():
			log.Printf("GRPC: %s disconnected from %s. Global context cancelled.", in.UserId, in.ServerId)
			return nil

		default:
			if err := srv.Context().Err(); err != nil {
				log.Println("GRPC: Server found the context to be done in the default case, cancelling global context")
				cancel()
				return nil
			}

			for message := range *s.queue.Messages {
				relay(message, srv, cancel)
			}
		}
	}
}

package server

import (
	"context"
	"fmt"
	"log"
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
}

func NewServer(cache *cache.Cache) *Server {
	return &Server{
		cache: cache,
	}
}

func (s *Server) Connect(in *pb.ConnectRequest, srv pb.ChatService_ConnectServer) error {
	// Create a unique streamid for this connection
	streamid := RandStringRunes(10)

	log.Printf("GRPC %s: user %s connected to server %s", streamid, in.UserId, in.ServerId)

	// Create a NATS queue for this streamid
	queue := queue.NewQueue("nats:4222", streamid)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go queue.Subscribe(ctx, in.ServerId)

	// send a "connected" message to the client to tell the client it successfully connected
	s.verifyConnection(srv, in)

	// getPastMessages
	if err := s.getPastMessages(srv, in); err != nil {
		log.Printf("error getting past messages: %v", err)
		return err
	}

	go ping(ctx, srv, cancel, streamid)
	// Receive messages from the NATS loop and forward them to the client
	for {
		select {
		case <-ctx.Done():
			log.Printf("GRPC %s: %s disconnected from %s. Global context cancelled.", streamid, in.UserId, in.ServerId)
			return nil

		default:
			if err := srv.Context().Err(); err != nil {
				log.Printf("GRPC %s: Server found the context to be done in the default case, cancelling global context", streamid)
				cancel()
				return nil
			}

			for message := range *queue.Messages {
				if err := relay(message, srv, cancel, streamid); err != nil {
					queue.ErrCh <- err
				}
			}
		}
	}
}

// Send receives a message from the client and publishes it to the NATS server
func (s *Server) Send(ctx context.Context, in *pb.ChatMessage) (*pb.SendResponse, error) {
	// log.Printf("GRPC: %s/%s->%s", in.UserId, in.ChannelId, in.Text)
	queue := queue.NewQueue("nats:4222", "NOT_USED")
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

	if err := queue.Publish("myserver", payload); err != nil {
		log.Printf("error publishing message to queue: %v", err)
		return nil, err
	}
	queue.Close()
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
func relay(message nats.Msg, srv pb.ChatService_ConnectServer, cancel context.CancelFunc, streamid string) error {
	// Convert JSON message to Notification object
	chatMsg, err := JSONToProto(message.Data)
	if err != nil {
		log.Printf("unmarshal error %v", err)
		return err
	}

	// Override the timestamp with the current time
	chatMsg.Ts = fmt.Sprint(time.Now().UnixNano())
	//log.Printf("N->G %s: %s %s %s: %s", streamid, chatMsg.ChannelId, chatMsg.UserId, chatMsg.Ts, chatMsg.Text)

	if err := srv.Send(chatMsg); err != nil {
		// If the client has disconnected, cancel the global context
		cancel()
		return err
	}

	return nil
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
			return err
		}
	}

	return nil
}

// Ping will send a ping message to the client every second and cancel the global context if the client disconnects
func ping(ctx context.Context, srv pb.ChatService_ConnectServer, cancel context.CancelFunc, streamid string) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			time.Sleep(1 * time.Second)

			err := srv.Send(&pb.ChatMessage{
				ChannelId: "system", // system information channel
				UserId:    "server",
				Text:      "ping",
				Ts:        "0",
			})

			if err != nil {
				cancel()
			}
		}
	}
}

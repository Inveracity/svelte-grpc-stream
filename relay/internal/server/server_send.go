package server

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/chat/v1"
	"github.com/inveracity/svelte-grpc-stream/internal/queue"
	"github.com/nats-io/nats.go"
)

// Send: receives messages from clients and publishes them to a NATS queue.
func (s *Server) Send(ctx context.Context, in *pb.ChatMessage) (*pb.SendResponse, error) {

	q := queue.NewQueue(s.natsURL, "")

	// Set message timestamp to the current time, this is done because the client may have a different time than the server.
	in.Ts = fmt.Sprint(time.Now().UnixNano())

	msg := nats.NewMsg(ServerID)

	payload, err := ProtoToJSON(in)
	if err != nil {
		return nil, err
	}

	msg.Data = payload

	if in.ChannelId != "system" { // only cache non-system messages
		if err := s.cache.Set(ServerID, string(payload)); err != nil {
			log.Printf("error writing message to cache: %v", err)
			return nil, err
		}
	}

	if err := q.Publish(ServerID, payload); err != nil {
		log.Printf("error publishing message to queue: %v", err)
		return nil, err
	}
	q.Close()
	return &pb.SendResponse{Ok: true, Error: ""}, nil
}

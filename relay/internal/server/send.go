package server

import (
	"context"
	"log"

	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/notifications/v1"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/encoding/protojson"
)

// Send receives a request from the client and publishes it to the NATS server
func (s *server) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	log.Printf("user: %s sent: %s to channel: %s", in.UserId, in.Text, in.ChannelId)
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
	msg.Data = []byte(payload)

	if err := nc.PublishMsg(msg); err != nil {
		return nil, err
	}

	return &pb.SendResponse{Ok: true, Error: ""}, nil
}

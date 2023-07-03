package server

import (
	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/chat/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func ProtoToJSON(in *pb.ChatMessage) ([]byte, error) {
	j := protojson.MarshalOptions{UseProtoNames: true}
	return j.Marshal(in)
}

func JSONToProto(in []byte) (*pb.ChatMessage, error) {
	var chatMsg pb.ChatMessage
	j := protojson.UnmarshalOptions{}
	if err := j.Unmarshal(in, &chatMsg); err != nil {
		return nil, err
	}
	return &chatMsg, nil
}

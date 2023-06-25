package server

import (
	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/notifications/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func ProtoToJSON(in *pb.Notification) ([]byte, error) {
	j := protojson.MarshalOptions{UseProtoNames: true}
	return j.Marshal(in)
}

func JSONToProto(in []byte) (*pb.Notification, error) {
	var notification pb.Notification
	j := protojson.UnmarshalOptions{}
	if err := j.Unmarshal(in, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

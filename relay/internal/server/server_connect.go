package server

import (
	"context"
	"log"
	"time"

	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/chat/v1"
	"github.com/inveracity/svelte-grpc-stream/internal/queue"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/encoding/protojson"
)

// Connect: receives a connection request from the client and creates a NATS queue subscriber
func (s *Server) Connect(in *pb.ConnectRequest, srv pb.ChatService_ConnectServer) error {
	// Create a unique streamid for this connection, this is used to trace logs for each connection
	s.streamid = RandStringRunes(10)

	log.Printf("GRPC %s: user %s connected to server %s", s.streamid, in.UserId, in.ServerId)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.queue = queue.NewQueue(s.natsURL, s.streamid)
	go s.queue.Subscribe(ctx, in.ServerId)

	// Send a "connected" message to the client to tell the client it successfully connected.
	srv.Send(systemMessage("connected", "server"))

	// Once the client is connected, send all past messages from the cache.
	if err := s.getHistory(srv, in); err != nil {
		log.Printf("error getting past messages: %v", err)
		return err
	}

	// Every client is continuously pinged to track whether they are still connected.
	// It also makes other users aware of who is connected to show in the UI.
	go s.ping(ctx, srv, in, cancel)

	// This is the main loop that receives messages from the NATS queue and sends them to the client.
	for {
		select {
		case <-ctx.Done():
			log.Printf("GRPC %s: %s disconnected from %s. Global context cancelled.", s.streamid, in.UserId, in.ServerId)
			return nil

		default:
			if err := srv.Context().Err(); err != nil {
				log.Printf("GRPC %s: Server found the context to be done in the default case, cancelling global context", s.streamid)
				cancel()
				return nil
			}

			for message := range *s.queue.Messages {
				if err := relay(message, srv, cancel, s.streamid); err != nil {
					s.queue.ErrCh <- err
				}
			}
		}
	}
}

// systemMessage is a convenience function to create a system message
func systemMessage(msg, userid string) *pb.ChatMessage {
	return &pb.ChatMessage{
		ChannelId: "system", // system information channel - the UI implements behavior based on events received on this channel
		UserId:    userid,   // 'server' is not an actual user
		Text:      msg,
		Ts:        "0",
	}
}

// GetHistory finds all the messages stored in Redis for a given ServerID
func (s *Server) getHistory(srv pb.ChatService_ConnectServer, in *pb.ConnectRequest) error {
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

// Send messages from NATS to the gRPC client
func relay(message nats.Msg, srv pb.ChatService_ConnectServer, cancel context.CancelFunc, streamid string) error {
	// Convert JSON message to Notification object
	chatMsg, err := JSONToProto(message.Data)
	if err != nil {
		log.Printf("unmarshal error %v", err)
		return err
	}

	if err := srv.Send(chatMsg); err != nil {
		// If the client has disconnected, cancel the global context. This closed the NATS queue and stops the main loop.
		// This should hopefully never happen as the ping function should cancel the global context if the client disconnects.
		cancel()
		return err
	}

	return nil
}

// Ping will send a ping message to the client every second and cancel the global context if the client disconnects
func (s *Server) ping(ctx context.Context, srv pb.ChatService_ConnectServer, in *pb.ConnectRequest, cancel context.CancelFunc) {
	for {
		select {
		case <-ctx.Done():
			err := s.broadcast(in.UserId, "disconnected")
			if err != nil {
				log.Printf("PING %s: error broadcasting disconnect message: %v", s.streamid, err)
			}
			return

		default:
			time.Sleep(1 * time.Second)
			s.broadcast(in.UserId, "connected")
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

// Broadcast sends a message to all connected clients
func (s *Server) broadcast(user, msg string) error {
	/*
		The channelid "system" is used to send system messages to the client that are not shown in the chat UI.
		The UI implements behavior based on events received on this channel.
		The "userId" field is used as metadata about which client the broadcast is coming from.
		The timestamp is an ignored field.
	*/
	return s.queue.Publish(ServerID, []byte(`{"channelId":"system","userId":"`+user+`","text":"`+msg+`","ts":"0"}`))
}

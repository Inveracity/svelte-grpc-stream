package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/inveracity/svelte-grpc-stream/internal/auth"
	"github.com/inveracity/svelte-grpc-stream/internal/cache"
	pb "github.com/inveracity/svelte-grpc-stream/internal/proto/chat/v1"
	"github.com/inveracity/svelte-grpc-stream/internal/queue"

	"google.golang.org/protobuf/encoding/protojson"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	cache    *cache.Cache
	queue    *queue.Queue
	streamid string
	natsURL  string
	pbURL    string
	pbAdmin  string
	pbPass   string
}

func NewServer(natsURL, pbURL, pbAdmin, pbPass string, cache *cache.Cache) *Server {
	return &Server{
		cache:   cache,
		natsURL: natsURL,
		pbURL:   pbURL,
		pbAdmin: pbAdmin,
		pbPass:  pbPass,
	}
}

func (s *Server) Connect(in *pb.ConnectRequest, srv pb.ChatService_ConnectServer) error {
	// Create a unique streamid for this connection
	s.streamid = RandStringRunes(10)

	log.Printf("GRPC %s: user %s connected to server %s", s.streamid, in.UserId, in.ServerId)

	auth := auth.New(s.pbURL, s.pbAdmin, s.pbPass)

	authed, err := auth.VerifyUserToken(in.Jwt)
	if err != nil {
		log.Printf("GRPC %s: error verifying jwt: %v", s.streamid, err)
		return fmt.Errorf("error verifying jwt")
	}

	if !authed {
		log.Printf("GRPC %s: user %s not authorized", s.streamid, in.UserId)
		return fmt.Errorf("user not authorized")
	}

	// Create a NATS queue subscriber for this s.streamid
	s.queue = queue.NewQueue(s.natsURL, s.streamid)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go s.queue.Subscribe(ctx, in.ServerId)

	// send a "connected" message to the client to tell the client it successfully connected
	srv.Send(systemMessage("connected", "server"))
	// getPastMessages
	if err := s.getPastMessages(srv, in); err != nil {
		log.Printf("error getting past messages: %v", err)
		return err
	}

	go s.ping(ctx, srv, in, cancel)
	// Receive messages from the NATS loop and forward them to the client
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

// Send: receives a message from the client and publishes it to the NATS server
func (s *Server) Send(ctx context.Context, in *pb.ChatMessage) (*pb.SendResponse, error) {

	auth := auth.New(s.pbURL, s.pbAdmin, s.pbPass)

	authed, err := auth.VerifyUserToken(in.Jwt)

	if err != nil || !authed {
		return nil, fmt.Errorf("user not authorized")
	}

	q := queue.NewQueue(s.natsURL, "")
	// Override timstamp
	in.Ts = fmt.Sprint(time.Now().UnixNano())

	msg := nats.NewMsg("myserver")

	payload, err := ProtoToJSON(in)
	if err != nil {
		return nil, err
	}

	msg.Data = payload

	if in.ChannelId != "system" { // only cache non-system messages
		if err := s.cache.Set("myserver", string(payload)); err != nil {
			log.Printf("error writing message to cache: %v", err)
			return nil, err
		}
	}

	if err := q.Publish("myserver", payload); err != nil {
		log.Printf("error publishing message to queue: %v", err)
		return nil, err
	}
	q.Close()
	return &pb.SendResponse{Ok: true, Error: ""}, nil
}

func systemMessage(msg, userid string) *pb.ChatMessage {
	return &pb.ChatMessage{
		ChannelId: "system", // system information channel - the UI implements behavior based on events received on this channel
		UserId:    userid,   // 'server' is not an actual user
		Text:      msg,
		Ts:        "0",
	}
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
	return s.queue.Publish("myserver", []byte(`{"channelId":"system","userId":"`+user+`","text":"`+msg+`","ts":"0"}`))
}

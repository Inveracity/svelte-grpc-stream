package server

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

func NatsSub(ctx context.Context, url, channelId string, events *chan nats.Msg) error {
	log.Printf("NATS: subscribing to channel: %s", channelId)
	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}

	subject := "events." + channelId
	msgChan := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe(subject, msgChan)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Printf("NATS: unsubscribing disconnected client: %s", channelId)
			sub.Unsubscribe()
			return nil

		default:
			msg, err := sub.NextMsgWithContext(ctx)
			if err != nil {
				log.Printf("next message error: %v", err)
				continue
			}

			// the nats message is sent back to the gRPC handler via the events channel, and will be "Ack()"ed there
			*events <- *msg
		}
	}
}

package queue

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

type Queue struct {
	nats     *nats.Conn
	Messages *chan nats.Msg
}

func NewQueue(nats *nats.Conn, Messages *chan nats.Msg) *Queue {
	return &Queue{
		nats:     nats,
		Messages: Messages,
	}
}

func (q *Queue) Publish(channel string, message []byte) error {
	return q.nats.Publish(channel, message)
}

func (q *Queue) Subscribe(ctx context.Context, channel string) error {
	msgChan := make(chan *nats.Msg, 64)
	sub, err := q.nats.ChanSubscribe(channel, msgChan)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			sub.Unsubscribe()
			log.Println("NATS: Ubsubbing because global context cancelled")
			return nil

		default:
			msg, err := sub.NextMsgWithContext(ctx)
			if err != nil {
				log.Printf("NATS: error getting next message from channel %s: %v", channel, err)
				continue
			}

			*q.Messages <- *msg
		}
	}

}

package queue

import (
	"context"

	"github.com/nats-io/nats.go"
)

type Queue struct {
	ctx      context.Context
	nats     *nats.Conn
	Messages *chan nats.Msg
}

func NewQueue(ctx context.Context, nats *nats.Conn, Messages *chan nats.Msg) *Queue {
	return &Queue{
		ctx:      ctx,
		nats:     nats,
		Messages: Messages,
	}
}

func (q *Queue) Publish(channel string, message []byte) error {
	return q.nats.Publish(channel, message)
}

func (q *Queue) Subscribe(channel string) error {
	msgChan := make(chan *nats.Msg, 64)
	sub, err := q.nats.ChanSubscribe(channel, msgChan)
	if err != nil {
		return err
	}

	for {
		select {
		case <-q.ctx.Done():
			sub.Unsubscribe()
			return nil
		default:
			msg, err := sub.NextMsgWithContext(q.ctx)
			if err != nil {
				continue
			}

			// the nats message is sent back to the gRPC handler
			// via the events channel, and will be "Ack()"ed there
			*q.Messages <- *msg
		}
	}

}

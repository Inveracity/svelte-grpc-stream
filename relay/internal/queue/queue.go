package queue

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Queue struct {
	nats     *nats.Conn
	Messages *chan nats.Msg
	ErrCh    chan error
	streamid string
}

func NewQueue(natsURL string, streamid string) *Queue {
	natsConn, err := nats.Connect(natsURL)
	if err != nil {
		panic(err)
	}

	errCh := make(chan error)
	messages := make(chan nats.Msg, 64)

	return &Queue{
		nats:     natsConn,
		Messages: &messages,
		ErrCh:    errCh,
		streamid: streamid,
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
		// Depending on where in the loop we are
		// either the context is cancelled or an error was caught on the error channel
		// Currently unsure why.
		case <-ctx.Done():
			time.Sleep(1 + time.Second) // TODO: There's a race condition when a client disconnects and the server wants to broadcast a message.
			sub.Unsubscribe()
			log.Printf("NATS %s: unsubscribing and closing stream", q.streamid)
			q.Close()
			return nil
		case <-q.ErrCh:
			sub.Unsubscribe()
			log.Printf("NATS %s: unsubscribing", q.streamid)
			return nil
		default:
			msg, err := sub.NextMsgWithContext(ctx)
			if err != nil {
				continue
			}

			*q.Messages <- *msg
		}
	}
}

func (q *Queue) Close() {
	q.nats.Close()
}

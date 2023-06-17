package queue

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func Client(ctx context.Context, url, subscriberId string, events *chan nats.Msg) error {
	log.Printf("NATS: subscriberId: %s", subscriberId)

	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}

	log.Print("connecting to jetstream")
	js, err := nc.JetStream()
	if err != nil {
		return err
	}

	cfg := &nats.StreamConfig{
		Name:      "EVENTS",
		Retention: nats.WorkQueuePolicy,
		Subjects:  []string{"events.>"},
	}

	js.AddStream(cfg)

	subject := "events." + subscriberId
	sub, err := js.PullSubscribe(subject, subscriberId, nats.BindStream(cfg.Name))
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Printf("unsubscribing disconnected client: %s", subscriberId)
			sub.Unsubscribe()
			return nil

		default:
			msgs, err := sub.Fetch(1, nats.MaxWait(1*time.Second))
			if err != nil {
				continue
			}
			if len(msgs) == 0 {
				continue
			}
			msg := msgs[0]
			// the nats message is sent back to the gRPC handler via the events channel, and will be "Ack()"ed there
			*events <- *msg
			log.Printf("Received a message: %s", string(msg.Data))
		}
	}
}

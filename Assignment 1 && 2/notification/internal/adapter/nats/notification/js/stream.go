package js

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func InitNotificationStream(nc *nats.Conn) {
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = js.CreateStream(ctx, jetstream.StreamConfig{
		Name:       "PAYMENTS",
		Subjects:   []string{"payment.completed"},
		Duplicates: 2 * time.Minute,
		MaxAge:     24 * time.Hour,
	})

	if err != nil {
		log.Fatal(err)
	}
}

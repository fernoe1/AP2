package js

import (
	"context"
	"encoding/json"
	"log"

	"github.com/fernoe1/AP2/assignment-1/notification/internal/adapter/nats/notification"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/domain"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/job"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NotificationConsumer struct {
	Worker *job.Worker
}

func (c *NotificationConsumer) ConsumeNotificationStream(nc *nats.Conn) {
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	cons, err := js.CreateOrUpdateConsumer(context.Background(), "PAYMENTS", jetstream.ConsumerConfig{
		Durable:       "NOTIFICATION_CONSUMER",
		FilterSubject: "payment.completed",
		MaxDeliver:    5,
	})

	if err != nil {
		log.Fatal(err)
	}

	msgs, err := cons.Messages()
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := msgs.Next()
		if err != nil {
			log.Println(err)
			continue
		}

		var notificationMsg notification.NotificationMessage
		if err := json.Unmarshal(msg.Data(), &notificationMsg); err != nil {
			log.Println(err)
			_ = msg.Ack()
			continue
		}

		err = c.Worker.Enqueue(context.Background(), domain.Notification{
			ID:            notificationMsg.OrderID,
			Amount:        notificationMsg.Amount,
			CustomerEmail: notificationMsg.CustomerEmail,
			Status:        notificationMsg.Status,
		})
		if err != nil {
			log.Printf("failed to enqueue notification: %v", err)
			continue
		}

		_ = msg.Ack()
	}
}

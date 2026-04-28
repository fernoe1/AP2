package js

import (
	"context"
	"encoding/json"
	"log"

	"github.com/fernoe1/AP2/assignment-1/notification/internal/adapter/nats/notification"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/domain"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/usecase"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NotificationConsumer struct {
	NotificationUsecase usecase.NotificationUsecase
	Processed           map[string]bool
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

		if c.Processed[notificationMsg.OrderID] {
			_ = msg.Ack()
			continue
		}

		c.NotificationUsecase.Send(context.Background(), &domain.Notification{
			ID:            notificationMsg.OrderID,
			Amount:        notificationMsg.Amount,
			CustomerEmail: notificationMsg.CustomerEmail,
			Status:        notificationMsg.Status,
		})

		c.Processed[notificationMsg.OrderID] = true

		_ = msg.Ack()
	}
}

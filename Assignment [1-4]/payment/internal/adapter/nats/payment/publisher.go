package payment

import (
	"context"
	"encoding/json"

	"github.com/fernoe1/AP2/assignment-1/payment/internal/domain"
	"github.com/nats-io/nats.go/jetstream"
)

type PaymentPublisher struct {
	Js jetstream.JetStream
}

func (pp *PaymentPublisher) PublishPaymentCompleted(ctx context.Context, payment *domain.Payment) error {
	msg := PaymentCompletedMessage{
		OrderID:       payment.OrderID,
		Amount:        payment.Amount,
		CustomerEmail: "krutoytemirlan2007@gmail.com",
		Status:        payment.Status,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = pp.Js.Publish(ctx, "payment.completed", data)
	if err != nil {
		return err
	}

	return nil
}

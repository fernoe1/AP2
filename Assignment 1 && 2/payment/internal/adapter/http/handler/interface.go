package handler

import (
	"context"

	"github.com/fernoe1/AP2/assignment-1/payment/internal/domain"
)

type PaymentUsecase interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) error
	GetPaymentFromOrderId(ctx context.Context, orderId string) ([]*domain.Payment, error)
}

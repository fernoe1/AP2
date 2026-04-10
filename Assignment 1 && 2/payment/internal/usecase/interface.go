package usecase

import (
	"context"

	"github.com/fernoe1/AP2/assignment-1/payment/internal/domain"
)

type PaymentRepository interface {
	SavePayment(ctx context.Context, payment *domain.Payment) error
	UpdatePayment(ctx context.Context, payment *domain.Payment) error
	FetchPaymentByOrderId(ctx context.Context, orderId string) ([]*domain.Payment, error)
}

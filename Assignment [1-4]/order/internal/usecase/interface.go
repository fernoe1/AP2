package usecase

import (
	"context"

	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, order *domain.Order) error
	FetchOrder(ctx context.Context, id uint) (*domain.Order, error)
	UpdateOrder(ctx context.Context, order *domain.Order) error
}

type OrderClient interface {
	GetOrderPaymentStatus(ctx context.Context, order *domain.Order) (string, error)
}

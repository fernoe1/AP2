package handler

import (
	"context"

	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
	GetOrder(ctx context.Context, id uint) (*domain.Order, error)
	CancelOrder(ctx context.Context, id uint) (*domain.Order, error)
	UpdateStatus(ctx context.Context, order *domain.Order, status string) error
}

package gorm

import (
	"context"

	NATS "github.com/fernoe1/AP2/assignment-1/order/internal/adapter/nats/order"
	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
	Nc *nats.Conn
}

func (r *OrderRepository) SaveOrder(ctx context.Context, order *domain.Order) error {
	return r.Db.WithContext(ctx).Create(order).Error
}

func (r *OrderRepository) FetchOrder(ctx context.Context, id uint) (*domain.Order, error) {
	var (
		order domain.Order
	)

	if err := r.Db.WithContext(ctx).First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) UpdateOrder(ctx context.Context, order *domain.Order) error {
	err := r.Db.WithContext(ctx).Save(order).Error
	if err != nil {
		return err
	}

	return NATS.PublishOrderUpdatedMessage(r.Nc, order)
}

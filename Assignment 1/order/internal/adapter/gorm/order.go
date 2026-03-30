package gorm

import (
	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func (r *OrderRepository) UpdateOrder(order *domain.Order) error {
	return r.Db.Save(order).Error
}

func (r *OrderRepository) FetchOrder(id uint) (*domain.Order, error) {
	var (
		order domain.Order
	)

	if err := r.Db.First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) SaveOrder(order *domain.Order) error {
	return r.Db.Create(order).Error
}

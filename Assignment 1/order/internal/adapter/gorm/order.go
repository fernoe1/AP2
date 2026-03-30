package gorm

import (
	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func (r *OrderRepository) UpdateOrder(order domain.Order) (*domain.Order, error) {
	result := r.Db.Save(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (r *OrderRepository) FetchOrder(id uint) (*domain.Order, error) {
	var (
		order domain.Order
	)

	result := r.Db.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (r *OrderRepository) SaveOrder(order domain.Order) error {
	return r.Db.Create(&order).Error
}

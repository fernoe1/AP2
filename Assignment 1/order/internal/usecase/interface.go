package usecase

import "github.com/fernoe1/AP2/assignment-1/order/internal/domain"

type OrderRepository interface {
	SaveOrder(order *domain.Order) error
	FetchOrder(id uint) (*domain.Order, error)
	UpdateOrder(order *domain.Order) error
}

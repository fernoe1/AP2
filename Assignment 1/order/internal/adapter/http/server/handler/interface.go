package handler

import "github.com/fernoe1/AP2/assignment-1/order/internal/domain"

type OrderUsecase interface {
	CreateOrder(order *domain.Order) error
	GetOrder(id uint) (*domain.Order, error)
	CancelOrder(id uint) (*domain.Order, error)
	UpdateStatus(order *domain.Order, status string) error
}

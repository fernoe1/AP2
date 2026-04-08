package client

import "github.com/fernoe1/AP2/assignment-1/order/internal/domain"

type OrderClient interface {
	GetOrderPaymentStatus(order *domain.Order) (string, error)
}

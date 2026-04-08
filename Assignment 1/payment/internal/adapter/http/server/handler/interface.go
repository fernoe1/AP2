package handler

import "github.com/fernoe1/AP2/assignment-1/payment/internal/domain"

type PaymentUsecase interface {
	CreatePayment(payment *domain.Payment) error
	GetPaymentFromOrderId(orderId string) ([]*domain.Payment, error)
}

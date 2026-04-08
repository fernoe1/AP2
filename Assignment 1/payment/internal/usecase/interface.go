package usecase

import "github.com/fernoe1/AP2/assignment-1/payment/internal/domain"

type PaymentRepository interface {
	SavePayment(payment *domain.Payment) error
	UpdatePayment(payment *domain.Payment) error
	FetchPaymentByOrderId(orderId string) ([]*domain.Payment, error)
}

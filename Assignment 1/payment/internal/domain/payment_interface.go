package domain

type PaymentUsecase interface {
	CreatePayment(payment *Payment) error
	GetPaymentFromOrderId(orderId string) ([]*Payment, error)
}

type PaymentRepository interface {
	SavePayment(payment *Payment) error
	UpdatePayment(payment *Payment) error
	FetchPaymentByOrderId(orderId string) ([]*Payment, error)
}

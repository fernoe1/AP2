package usecase

import (
	"fmt"

	"github.com/fernoe1/AP2/assignment-1/payment/internal/domain"
)

type PaymentUsecase struct {
	PaymentRepository domain.PaymentRepository
}

func (uc *PaymentUsecase) GetPaymentFromOrderId(orderId string) ([]*domain.Payment, error) {
	return uc.PaymentRepository.FetchPaymentByOrderId(orderId)
}

func (uc *PaymentUsecase) CreatePayment(payment *domain.Payment) error {
	status := "Authorized"
	if payment.Amount > 100000 {
		status = "Declined"
	}
	payment.Status = status

	if err := uc.PaymentRepository.SavePayment(payment); err != nil {
		return err
	}

	payment.TransactionID = domain.GenerateTransactionID(fmt.Sprintf("%d%s%d%s",
		payment.ID,
		payment.OrderID,
		payment.Amount,
		payment.Status,
	),
	)

	if err := uc.PaymentRepository.UpdatePayment(payment); err != nil {
		return err
	}

	return nil
}

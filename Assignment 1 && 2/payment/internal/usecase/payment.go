package usecase

import (
	"context"
	"fmt"

	"github.com/fernoe1/AP2/assignment-1/payment/internal/domain"
)

type PaymentUsecase struct {
	PaymentRepository PaymentRepository
	PaymentPublisher  PaymentPublisher
}

func (uc *PaymentUsecase) GetPaymentFromOrderId(ctx context.Context, orderId string) ([]*domain.Payment, error) {
	return uc.PaymentRepository.FetchPaymentByOrderId(ctx, orderId)
}

func (uc *PaymentUsecase) CreatePayment(ctx context.Context, payment *domain.Payment) error {
	status := "Authorized"
	if payment.Amount > 100000 {
		status = "Declined"
	}
	payment.Status = status

	if err := uc.PaymentRepository.SavePayment(ctx, payment); err != nil {
		return err
	}

	payment.TransactionID = domain.GenerateTransactionID(fmt.Sprintf("%d%s%d%s",
		payment.ID,
		payment.OrderID,
		payment.Amount,
		payment.Status,
	),
	)

	if err := uc.PaymentRepository.UpdatePayment(ctx, payment); err != nil {
		return err
	}

	return uc.PaymentPublisher.PublishPaymentCompleted(ctx, payment)
}

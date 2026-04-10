package gorm

import (
	"context"

	"github.com/fernoe1/AP2/assignment-1/payment/internal/domain"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	Db *gorm.DB
}

func (r *PaymentRepository) FetchPaymentByOrderId(ctx context.Context, orderId string) ([]*domain.Payment, error) {
	var (
		payment []*domain.Payment
	)

	if err := r.Db.WithContext(ctx).Where("order_id = ?", orderId).Find(&payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *PaymentRepository) UpdatePayment(ctx context.Context, payment *domain.Payment) error {
	return r.Db.WithContext(ctx).Save(&payment).Error
}

func (r *PaymentRepository) SavePayment(ctx context.Context, payment *domain.Payment) error {
	return r.Db.WithContext(ctx).Create(&payment).Error
}

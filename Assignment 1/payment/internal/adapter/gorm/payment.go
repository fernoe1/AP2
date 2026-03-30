package gorm

import (
	"github.com/fernoe1/AP2/assignment-1/payment/internal/domain"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	Db *gorm.DB
}

func (r *PaymentRepository) FetchPaymentByOrderId(orderId string) ([]*domain.Payment, error) {
	var (
		payment []*domain.Payment
	)

	if err := r.Db.Where("order_id = ?", orderId).Find(&payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *PaymentRepository) UpdatePayment(payment *domain.Payment) error {
	return r.Db.Save(&payment).Error
}

func (r *PaymentRepository) SavePayment(payment *domain.Payment) error {
	return r.Db.Create(&payment).Error
}

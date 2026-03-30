package usecase

import (
	"errors"

	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
)

type OrderUsecase struct {
	OrderRepository domain.OrderRepository
}

func (uc *OrderUsecase) CancelOrder(id uint) (*domain.Order, error) {
	order, err := uc.GetOrder(id)
	if err != nil {
		return nil, err
	}

	if order.Status != "Pending" {
		return nil, errors.New("order is not pending")
	}

	order.Status = "Cancelled"

	if err := uc.OrderRepository.UpdateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (uc *OrderUsecase) GetOrder(id uint) (*domain.Order, error) {
	return uc.OrderRepository.FetchOrder(id)
}

func (uc *OrderUsecase) CreateOrder(order *domain.Order) error {
	order.Status = "Pending"

	return uc.OrderRepository.SaveOrder(order)
}

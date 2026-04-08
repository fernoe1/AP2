package usecase

import (
	"errors"

	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http/client"
	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
)

type OrderUsecase struct {
	OrderRepository OrderRepository
	OrderClient     client.OrderClient
}

func (uc *OrderUsecase) UpdateStatus(order *domain.Order, status string) error {
	if status == "Authorized" {
		order.Status = "Paid"
	} else {
		order.Status = "Failed"
	}

	return uc.OrderRepository.UpdateOrder(order)
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

	if err := uc.OrderRepository.SaveOrder(order); err != nil {
		return err
	}

	status, err := uc.OrderClient.GetOrderPaymentStatus(order)
	if err != nil {
		return err
	}

	return uc.UpdateStatus(order, status)
}

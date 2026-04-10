package usecase

import (
	"context"
	"errors"

	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
)

type OrderUsecase struct {
	OrderRepository OrderRepository
	OrderClient     OrderClient
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, order *domain.Order) error {
	order.Status = "Pending"

	if err := uc.OrderRepository.SaveOrder(ctx, order); err != nil {
		return err
	}

	status, err := uc.OrderClient.GetOrderPaymentStatus(ctx, order)
	if err != nil {
		return err
	}

	return uc.UpdateStatus(ctx, order, status)
}

func (uc *OrderUsecase) GetOrder(ctx context.Context, id uint) (*domain.Order, error) {
	return uc.OrderRepository.FetchOrder(ctx, id)
}

func (uc *OrderUsecase) CancelOrder(ctx context.Context, id uint) (*domain.Order, error) {
	order, err := uc.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	if order.Status != "Pending" {
		return nil, errors.New("order is not pending")
	}

	order.Status = "Cancelled"

	if err := uc.OrderRepository.UpdateOrder(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (uc *OrderUsecase) UpdateStatus(ctx context.Context, order *domain.Order, status string) error {
	if status == "Authorized" {
		order.Status = "Paid"
	} else {
		order.Status = "Failed"
	}

	return uc.OrderRepository.UpdateOrder(ctx, order)
}

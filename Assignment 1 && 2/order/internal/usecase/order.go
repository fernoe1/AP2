package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
	"github.com/redis/go-redis/v9"
)

type OrderUsecase struct {
	OrderRepository OrderRepository
	OrderClient     OrderClient
	RDB             *redis.Client
	TTL             time.Duration
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
	key := strconv.Itoa(int(id))

	cached, err := uc.RDB.Get(ctx, key).Result()

	if err == nil {
		var order domain.Order
		if err := json.Unmarshal([]byte(cached), &order); err == nil {
			return &order, nil
		}
	}

	if !errors.Is(err, redis.Nil) && err != nil {
		fmt.Printf("redis get failed: %v\n", err)
	}

	order, err := uc.OrderRepository.FetchOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	raw, err := json.Marshal(&order)
	if err == nil {
		_ = uc.RDB.Set(ctx, key, raw, uc.TTL).Err()
	}

	return order, nil
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

	if err := uc.OrderRepository.UpdateOrder(ctx, order); err != nil {
		return err
	}

	if err := uc.RDB.Del(ctx, strconv.Itoa(int(order.ID))).Err(); err != nil {
		fmt.Printf("redis del failed: %v\n", err)
	}

	return nil
}

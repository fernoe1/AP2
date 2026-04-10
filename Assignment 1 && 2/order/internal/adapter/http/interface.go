package http

import (
	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http/handler"
)

type OrderUsecase interface {
	handler.OrderUsecase
}

package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	client "github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http"
	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http/dto"
	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
	"github.com/fernoe1/AP2/assignment-1/order/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler struct {
	OrderUsecase domain.OrderUsecase
}

func (h *OrderHandler) Patch(c *gin.Context) {
	var (
		rb = util.ResponseBuilder{C: c}
	)

	id := c.Param("id")
	uId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		rb.Response(http.StatusBadRequest, err.Error(), nil)

		return
	}

	order, err := h.OrderUsecase.CancelOrder(uint(uId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			rb.Response(http.StatusNotFound, err.Error(), nil)

			return
		}

		rb.Response(http.StatusBadRequest, err.Error(), nil)

		return
	}

	rb.Response(http.StatusOK, "ok", order)
}

func (h *OrderHandler) Post(c *gin.Context) {
	var (
		opDTO dto.Order
		rb    = util.ResponseBuilder{C: c}
	)

	if err := c.ShouldBindJSON(&opDTO); err != nil {
		rb.Response(http.StatusBadRequest, err.Error(), nil)

		return
	}

	order := domain.Order{
		CustomerID: opDTO.CustomerID,
		ItemName:   opDTO.ItemName,
		Amount:     opDTO.Amount,
	}

	if err := h.OrderUsecase.CreateOrder(&order); err != nil {
		rb.Response(http.StatusInternalServerError, err.Error(), nil)

		return
	}

	payment, _ := json.Marshal(dto.Payment{
		OrderID: strconv.Itoa(int(order.ID)),
		Amount:  order.Amount},
	)

	resp, err := client.C.Post(
		"http://localhost:8082/payments",
		"application/json",
		bytes.NewBuffer(payment),
	)

	if err != nil {
		rb.Response(http.StatusServiceUnavailable, err.Error(), nil)

		return
	}

	var data dto.Data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		rb.Response(http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if err := h.OrderUsecase.UpdateStatus(&order, data.Data["Status"].(string)); err != nil {
		rb.Response(http.StatusInternalServerError, err.Error(), nil)
	}

	rb.Response(http.StatusCreated, "ok", order)
}

func (h *OrderHandler) Get(c *gin.Context) {
	var (
		rb = util.ResponseBuilder{C: c}
	)

	id := c.Param("id")
	uId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		rb.Response(http.StatusBadRequest, err.Error(), nil)

		return
	}

	order, err := h.OrderUsecase.GetOrder(uint(uId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			rb.Response(http.StatusNotFound, err.Error(), nil)

			return
		}

		rb.Response(http.StatusInternalServerError, err.Error(), nil)

		return
	}

	rb.Response(http.StatusOK, "ok", order)
}

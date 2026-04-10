package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http/dto"
	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
	"github.com/fernoe1/AP2/assignment-1/order/pkg/app"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler struct {
	OrderUsecase OrderUsecase
}

func (h *OrderHandler) PatchStatus(c *gin.Context) {
	var (
		patchSQ dto.PatchStatusRequest
		rb      = app.ResponseBuilder{C: c}
	)

	if err := c.ShouldBindJSON(&patchSQ); err != nil {
		rb.Response(http.StatusBadRequest, err.Error(), nil)

		return
	}

	id := c.Param("id")
	uId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		rb.Response(http.StatusBadRequest, err.Error(), nil)

		return
	}

	order := &domain.Order{
		ID:     uint(uId),
		Status: patchSQ.Status,
	}

	if err := h.OrderUsecase.UpdateStatus(c, order, order.Status); err != nil {
		rb.Response(http.StatusInternalServerError, err.Error(), nil)
	}

	rb.Response(http.StatusOK, "ok", order)
}

func (h *OrderHandler) Patch(c *gin.Context) {
	var (
		rb = app.ResponseBuilder{C: c}
	)

	id := c.Param("id")
	uId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		rb.Response(http.StatusBadRequest, err.Error(), nil)

		return
	}

	order, err := h.OrderUsecase.CancelOrder(c, uint(uId))
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
		rb    = app.ResponseBuilder{C: c}
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

	if err := h.OrderUsecase.CreateOrder(c, &order); err != nil {
		rb.Response(http.StatusInternalServerError, err.Error(), nil)

		return
	}

	rb.Response(http.StatusCreated, "ok", order)
}

func (h *OrderHandler) Get(c *gin.Context) {
	var (
		rb = app.ResponseBuilder{C: c}
	)

	id := c.Param("id")
	uId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		rb.Response(http.StatusBadRequest, err.Error(), nil)

		return
	}

	order, err := h.OrderUsecase.GetOrder(c, uint(uId))
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

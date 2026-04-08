package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http/dto"
	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
	"github.com/fernoe1/AP2/assignment-1/order/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler struct {
	OrderUsecase OrderUsecase
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

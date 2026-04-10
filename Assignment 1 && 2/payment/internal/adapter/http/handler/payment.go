package handler

import (
	"errors"
	"net/http"

	"github.com/fernoe1/AP2/assignment-1/payment/internal/adapter/http/dto"
	"github.com/fernoe1/AP2/assignment-1/payment/internal/domain"
	"github.com/fernoe1/AP2/assignment-1/payment/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaymentHandler struct {
	PaymentUsecase PaymentUsecase
}

func (h *PaymentHandler) Get(c *gin.Context) {
	var (
		rb = util.ResponseBuilder{C: c}
	)

	orderId := c.Param("order_id")

	payment, err := h.PaymentUsecase.GetPaymentFromOrderId(c, orderId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			rb.Response(http.StatusNotFound, err.Error(), nil)

			return
		}

		rb.Response(http.StatusInternalServerError, err.Error(), nil)

		return
	}

	rb.Response(http.StatusOK, "ok", payment)
}

func (h *PaymentHandler) Post(c *gin.Context) {
	var (
		ppDTO dto.PaymentPostDTO
		rb    = util.ResponseBuilder{C: c}
	)

	if err := c.ShouldBindJSON(&ppDTO); err != nil {
		rb.Response(http.StatusBadRequest, err.Error(), nil)

		return
	}

	payment := domain.Payment{
		OrderID: ppDTO.OrderID,
		Amount:  ppDTO.Amount,
	}

	if err := h.PaymentUsecase.CreatePayment(c, &payment); err != nil {
		rb.Response(http.StatusInternalServerError, err.Error(), nil)

		return
	}

	rb.Response(http.StatusCreated, "ok", payment)
}

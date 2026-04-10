package http

import "C"
import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http/dto"
	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
)

type Client struct {
	C *http.Client
}

func InitClient() *Client {
	return &Client{C: &http.Client{Timeout: 2 * time.Second}}
}

func (c *Client) GetOrderPaymentStatus(ctx context.Context, order *domain.Order) (string, error) {
	paymentDto := dto.Payment{
		OrderID: strconv.Itoa(int(order.ID)),
		Amount:  order.Amount,
	}

	paymentJson, _ := json.Marshal(paymentDto)

	resp, err := c.C.Post(
		"http://payment:8082/payments",
		"application/json",
		bytes.NewBuffer(paymentJson),
	)

	if err != nil {
		return "", err
	}

	var data dto.Data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.Data["Status"].(string), nil
}

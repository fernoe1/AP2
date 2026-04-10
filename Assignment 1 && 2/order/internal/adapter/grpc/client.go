package grpc

import (
	"context"
	"strconv"

	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
	svc "github.com/fernoe1/protogen/ap2-assign2/service/payment"
)

type Client struct {
	C svc.PaymentServiceClient
}

func InitClient(client svc.PaymentServiceClient) *Client {
	return &Client{C: client}
}

func (c *Client) GetOrderPaymentStatus(ctx context.Context, order *domain.Order) (string, error) {
	createResp, err := c.C.CreatePayment(ctx, &svc.CreatePaymentRequest{
		OrderId: strconv.Itoa(int(order.ID)),
		Amount:  order.Amount,
	})

	if err != nil {
		return "", err
	}

	return createResp.Payment.Status, nil
}

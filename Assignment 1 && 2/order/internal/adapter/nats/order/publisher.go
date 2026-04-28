package order

import (
	"encoding/json"
	"strconv"

	"github.com/fernoe1/AP2/assignment-1/order/internal/domain"
	"github.com/nats-io/nats.go"
)

func PublishOrderUpdatedMessage(nc *nats.Conn, order *domain.Order) error {
	msg := OrderUpdatedMessage{
		OrderID: strconv.Itoa(int(order.ID)),
		Status:  order.Status,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	subj := "orders.updated." + strconv.Itoa(int(order.ID))

	return nc.Publish(subj, data)
}

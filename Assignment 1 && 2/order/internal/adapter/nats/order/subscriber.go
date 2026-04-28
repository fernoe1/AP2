package order

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
)

func SubscribeToOrderUpdates(
	nc *nats.Conn, orderID int,
	handler func(OrderUpdatedMessage) error,
) error {
	subj := "orders.updated." + strconv.Itoa(orderID)

	sub, err := nc.SubscribeSync(subj)
	if err != nil {
		return err
	}

	defer sub.Unsubscribe()

	for {
		msg, err := sub.NextMsg(1 * time.Minute)
		if err != nil {
			return err
		}

		var event OrderUpdatedMessage
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			continue
		}

		if err := handler(event); err != nil {
			return err
		}
	}
}

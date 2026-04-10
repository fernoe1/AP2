package service

import (
	"encoding/json"
	"strconv"
	"time"

	NATS "github.com/fernoe1/AP2/assignment-1/order/internal/adapter/nats"
	svc "github.com/fernoe1/protogen/ap2-assign2/service/order"
	"github.com/nats-io/nats.go"
)

type OrderService struct {
	svc.UnimplementedOrderServiceServer
	Nc *nats.Conn
}

func (o *OrderService) SubscribeToOrderUpdates(request *svc.OrderRequest, server svc.OrderService_SubscribeToOrderUpdatesServer) error {
	subj := "orders.updated." + strconv.Itoa(int(request.Id))

	sub, err := o.Nc.SubscribeSync(subj)
	if err != nil {
		return err
	}

	defer sub.Unsubscribe()

	for {
		msg, err := sub.NextMsg(1 * time.Minute)
		if err != nil {
			return err
		}

		var event NATS.OrderUpdatedEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			continue
		}

		err = server.Send(&svc.OrderStatusUpdate{
			Status: event.Status,
		})
		if err != nil {
			return err
		}
	}
}

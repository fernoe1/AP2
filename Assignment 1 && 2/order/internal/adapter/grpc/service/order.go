package service

import (
	NATS "github.com/fernoe1/AP2/assignment-1/order/internal/adapter/nats/order"
	svc "github.com/fernoe1/protogen/ap2-assign2/service/order"
	"github.com/nats-io/nats.go"
)

type OrderService struct {
	svc.UnimplementedOrderServiceServer
	Nc *nats.Conn
}

func (o *OrderService) SubscribeToOrderUpdates(request *svc.OrderRequest, server svc.OrderService_SubscribeToOrderUpdatesServer) error {
	return NATS.SubscribeToOrderUpdates(
		o.Nc,
		int(request.Id),
		func(event NATS.OrderUpdatedMessage) error {
			return server.Send(&svc.OrderStatusUpdate{
				Status: event.Status,
			})
		},
	)
}

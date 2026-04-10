package nats

type OrderUpdatedEvent struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

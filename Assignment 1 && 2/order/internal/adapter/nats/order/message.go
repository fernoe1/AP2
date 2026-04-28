package order

type OrderUpdatedMessage struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

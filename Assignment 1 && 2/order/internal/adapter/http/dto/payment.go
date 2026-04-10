package dto

type Payment struct {
	OrderID string `json:"order_id"`
	Amount  int64  `json:"amount"`
}

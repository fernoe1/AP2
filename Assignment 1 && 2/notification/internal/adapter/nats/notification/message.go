package notification

type NotificationMessage struct {
	OrderID       uint   `json:"order_id"`
	Amount        int64  `json:"amount"`
	CustomerEmail string `json:"customer_email"`
	Status        string `json:"status"`
}

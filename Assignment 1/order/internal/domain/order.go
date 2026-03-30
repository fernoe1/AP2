package domain

import "time"

type Order struct {
	ID         uint
	CustomerID string
	ItemName   string
	Amount     int64  // amount in cents
	Status     string // Pending/Paid/Cancelled/Failed
	CreatedAt  time.Time
}

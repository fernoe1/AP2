package dto

type PaymentPostDTO struct {
	OrderID string `json:"order_id" binding:"required"`
	Amount  int64  `json:"amount" binding:"required,min=1"`
}

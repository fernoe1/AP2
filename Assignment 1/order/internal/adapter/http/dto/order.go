package dto

type Order struct {
	CustomerID string `json:"customer_id" binding:"required"`
	ItemName   string `json:"item_name" binding:"required"`
	Amount     int64  `json:"amount" binding:"required,min=1"`
}

type PatchStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

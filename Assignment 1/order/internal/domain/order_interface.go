package domain

type OrderUsecase interface {
	CreateOrder(order Order) error
	GetOrder(id uint) (*Order, error)
	CancelOrder(id uint) (*Order, error)
}

type OrderRepository interface {
	SaveOrder(order Order) error
	FetchOrder(id uint) (*Order, error)
	UpdateOrder(order Order) (*Order, error)
}

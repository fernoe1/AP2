package domain

type Notification struct {
	ID            uint
	Amount        int64
	CustomerEmail string
	Status        string
}

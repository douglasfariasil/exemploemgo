package entity

type OrderRepositoryInterface interface {
	Save(Order *Order) error
	GetTotalTransact() (int, error)
}
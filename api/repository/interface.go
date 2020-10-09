package repository

//go:generate mockgen -package mock -source interface.go -destination ./mock/repository.go

type CustomerRepository interface {
	DeleteByPrefix(prefix []string) (customers []CustomerInfo, err error)
	GetByPrefix(prefix []string) ([]*CustomerInfo, error)
	CreateCustomer(customer CustomerInfo) (result int, err error)
}

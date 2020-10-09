package repository

import "context"

//go:generate mockgen -package mock -source interface.go -destination ./mock/repository.go

type CustomerRepository interface {
	DeleteByPrefix(ctx context.Context, prefix []string) (customers []CustomerInfo, err error)
	GetByPrefix(ctx context.Context, prefix []string) ([]*CustomerInfo, error)
	CreateCustomer(ctx context.Context, customer CustomerInfo) (result int, err error)
}

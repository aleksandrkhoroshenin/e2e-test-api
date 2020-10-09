package service

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/cobu/backend/e2e-test-api/api/repository"
)

type CustomerService struct {
	customerRepository repository.CustomerRepository
	logger             *logrus.Logger
}

func NewCustomerService(customerRepository repository.CustomerRepository, logger *logrus.Logger) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
		logger:             logger,
	}
}

func (cs *CustomerService) DeleteCustomers(ctx context.Context, prefixes []string) (*repository.DeletedCustomersInfo, error) {
	if len(prefixes) == 0 {
		cs.logger.Info("prefixes is empty")

		return nil, nil
	}

	customers, err := cs.customerRepository.DeleteByPrefix(ctx, prefixes)
	if err != nil {
		return nil, err
	}

	ids := make([]int, 0)
	for _, customer := range customers {
		ids = append(ids, customer.ID)
	}

	return &repository.DeletedCustomersInfo{
		Count: len(ids),
		IDs:   ids,
	}, nil
}

func (cs *CustomerService) GetCustomers(ctx context.Context, prefixes []string) ([]*repository.CustomerInfo, error) {
	if len(prefixes) == 0 {
		cs.logger.Info("prefixes is empty")

		return nil, nil
	}

	return cs.customerRepository.GetByPrefix(ctx, prefixes)
}

func (cs *CustomerService) CreateCustomer(ctx context.Context, customer repository.CustomerInfo) (int, error) {
	return cs.customerRepository.CreateCustomer(ctx, customer)
}

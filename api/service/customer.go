package service

import (
	"github.com/sirupsen/logrus"

	"github.com/cobu/backend/e2e-test-api/api/repository"
)

type CustomerService struct {
	customerRepository repository.CustomerRepository
	logger *logrus.Logger
}

func NewCustomerService(customerRepository repository.CustomerRepository, logger *logrus.Logger) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
		logger: logger,
	}
}

func (cs *CustomerService) DeleteCustomers(prefixes []string) (*repository.DeletedCustomersInfo, error) {
	if len(prefixes) == 0 {
		cs.logger.Info("prefixes is empty")

		return nil, nil
	}

	customers, err := cs.customerRepository.DeleteByPrefix(prefixes)
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

func (cs *CustomerService) GetCustomers(prefixes []string) ([]*repository.CustomerInfo, error) {
	if len(prefixes) == 0 {
		cs.logger.Info("prefixes is empty")

		return nil, nil
	}

	return cs.customerRepository.GetByPrefix(prefixes)
}

func (cs *CustomerService) CreateCustomer(customer repository.CustomerInfo) (int, error) {
	return cs.customerRepository.CreateCustomer(customer)
}
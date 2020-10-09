package api

import (
	"github.com/go-resty/resty/v2"
)

type Client interface {
	CreateCustomer(req CreateCustomerReq) (*resty.Response, error)
	ListCustomerByPrefix(req CustomerPrefixesReq) (*resty.Response, error)
	DeleteCustomerByPrefix(req CustomerPrefixesReq) (*resty.Response, error)
}

type client struct {
	client   *resty.Client
	endpoint string
}

func New(endpoint string) Client {
	return client{
		client:   resty.New(),
		endpoint: endpoint,
	}
}

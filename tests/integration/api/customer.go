package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (c client) CreateCustomer(req CreateCustomerReq) (*resty.Response, error){
	return c.client.R().
		SetBody(req).
		Post(fmt.Sprintf("%v%v", c.endpoint, "/api/v1/customer/createCustomer"))
}

func (c client) ListCustomerByPrefix(req CustomerPrefixesReq) (*resty.Response, error){
	return c.client.R().
		SetBody(req).
		Post(fmt.Sprintf("%v%v", c.endpoint, "/api/v1/customer/getCustomersByPrefix"))
}

func (c client) DeleteCustomerByPrefix(req CustomerPrefixesReq) (*resty.Response, error){
	return c.client.R().
		SetBody(req).
		Post(fmt.Sprintf("%v%v", c.endpoint, "/api/v1/customer/deleteCustomersByPrefix"))
}


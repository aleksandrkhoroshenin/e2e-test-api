package controller

import (
	"github.com/cobu/backend/e2e-test-api/api/handlers"
	"github.com/cobu/backend/e2e-test-api/api/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CustomerController struct {
	customerHandler *handlers.CustomerHandler
}

func NewCustomerController(customerHandler *service.CustomerService, logger *logrus.Logger) *CustomerController {
	return &CustomerController{customerHandler:
		handlers.NewCustomerHandler(customerHandler, logger)}
}

func (cc *CustomerController) RestController() *mux.Router {
	routers := mux.NewRouter().StrictSlash(true)

	routers.HandleFunc("/api/v1/customer/getCustomersByPrefix", cc.customerHandler.HandleListCustomerByPrefix).Methods(http.MethodPost)
	routers.HandleFunc("/api/v1/customer/deleteCustomersByPrefix", cc.customerHandler.HandleDeleteCustomerByPrefix).Methods(http.MethodPost)
	routers.HandleFunc("/api/v1/customer/createCustomer", cc.customerHandler.HandleCreateCustomer).Methods(http.MethodPost)

	return routers
}

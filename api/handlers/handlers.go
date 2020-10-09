package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/cobu/backend/e2e-test-api/api/repository"
	"github.com/cobu/backend/e2e-test-api/api/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CustomerHandler struct {
	customerService *service.CustomerService
	logger          *logrus.Logger
}

func NewCustomerHandler(customerService *service.CustomerService, logger *logrus.Logger) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
		logger:          logger,
	}
}

func (ch *CustomerHandler) HandleDeleteCustomerByPrefix(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	prefixes := CustomerPrefixesReq{}
	err = json.NewDecoder(r.Body).Decode(&prefixes)
	if err != nil {
		ch.logger.Errorf("decode body error: %v", err)

		return
	}

	info, err := ch.customerService.DeleteCustomers(prefixes.Prefixes)
	if err != nil {
		ch.logger.Errorf("delete customers error: %v", err)

		return
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(convertDeletedCustomersInfoResp(info))
	if err != nil {
		ch.logger.Errorf("encode error: %v", err)

		return
	}

	ch.logger.Info("success delete customers by prefixes")

	if _, err = w.Write(buf.Bytes()); err != nil {
		ch.logger.Errorf("write err: %v", err)
	}
}

func (ch *CustomerHandler) HandleListCustomerByPrefix(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	prefixes := CustomerPrefixesReq{}
	err = json.NewDecoder(r.Body).Decode(&prefixes)
	if err != nil {
		ch.logger.Errorf("decode body error: %v", err)

		return
	}

	customersInfo, err := ch.customerService.GetCustomers(prefixes.Prefixes)
	if err != nil {
		ch.logger.Errorf("get customers error: %v", err)

		return
	}

	if len(customersInfo) == 0 {
		ch.logger.Error("len customersInfo is zero")
		w.WriteHeader(http.StatusNotFound)

		return
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(convertCustomerInfo(customersInfo))
	if err != nil {
		ch.logger.Errorf("encode error: %v", err)

		return
	}

	ch.logger.Info("success list by prefixes")

	if _, err = w.Write(buf.Bytes()); err != nil {
		ch.logger.Errorf("write err: %v", err)
	}
}

func (ch *CustomerHandler) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	customerReq := CreateCustomerReq{}
	err = json.NewDecoder(r.Body).Decode(&customerReq)
	if err != nil {
		ch.logger.Errorf("decode body error: %v", err)

		return
	}

	count, err := ch.customerService.CreateCustomer(convertCreateCustomerReq(customerReq))
	if err != nil {
		ch.logger.Errorf("create customers error: %v", err)

		return
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(convertCreateCustomerResp(count))
	if err != nil {
		ch.logger.Errorf("encode error: %v", err)

		return
	}

	ch.logger.Info("success create customer")

	if _, err = w.Write(buf.Bytes()); err != nil {
		ch.logger.Errorf("write err: %v", err)
	}
}

func convertCustomerInfo(infos []*repository.CustomerInfo) (resp []CustomerResp) {
	for _, info := range infos {
		resp = append(resp, CustomerResp{
			FirstName:      info.FirstName,
			LastName:       info.LastName,
			PatronymicName: info.PatronymicName,
			Phone:          info.Phone,
			Email:          info.Email,
			ID:             info.ID,
		})
	}

	return resp
}

func convertCreateCustomerReq(req CreateCustomerReq) (customer repository.CustomerInfo) {
	return repository.CustomerInfo{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		PatronymicName: req.PatronymicName,
		Phone:          req.Phone,
		Email:          req.Email,
		ID:             req.ID,
	}
}

func convertCreateCustomerResp(count int) CreateCustomerResp {
	return CreateCustomerResp{
		CountInserted: count,
	}
}

func convertDeletedCustomersInfoResp(info *repository.DeletedCustomersInfo) DeleteCustomersResp {
	return DeleteCustomersResp{
		Count: info.Count,
		IDs:   info.IDs,
	}
}
package integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/stretchr/testify/assert"

	"github.com/cobu/backend/e2e-test-api/tests/integration/api"
)

const endpoint = "http://localhost:8090"

func TestCustomerCRUD(t *testing.T) {
	t.Run("delete created customer by prefix", func(t *testing.T) {
		customerID := 1488

		apiClient := api.New(endpoint)

		customerReq := api.CreateCustomerReq{}
		gofakeit.Struct(&customerReq)

		customerReq.ID = customerID

		resp, err := apiClient.CreateCustomer(customerReq)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode())

		var customerResp api.CreateCustomerResp

		err = json.Unmarshal(resp.Body(), &customerResp)
		assert.NoError(t, err)

		assert.Equal(t, 1, customerResp.CountInserted)

		resp, err = apiClient.DeleteCustomerByPrefix(api.CustomerPrefixesReq{
			Prefixes: []string{(*customerReq.FirstName)[0:3]},
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode())

		var deleteCustomerResp api.DeleteCustomersResp

		err = json.Unmarshal(resp.Body(), &deleteCustomerResp)
		assert.NoError(t, err)

		assert.Equal(t, 1, deleteCustomerResp.Count)

		assert.Equal(t, customerID, deleteCustomerResp.IDs[0])
	})
}

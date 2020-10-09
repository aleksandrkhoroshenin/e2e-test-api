package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/sirupsen/logrus"

	"github.com/cobu/backend/e2e-test-api/api/controller"
	"github.com/cobu/backend/e2e-test-api/api/repository"
	"github.com/cobu/backend/e2e-test-api/api/service"
)

func main() {
	servicePort := "8090"
	log.Printf("REST API started at %s...\n", servicePort)

	dbConnectionUrl, exists := os.LookupEnv("DB_CONNECTION_URL")
	if !exists {
		log.Fatal("DB_CONNECTION_URL env variable does not exist")
	}

	options, err := pg.ParseURL(dbConnectionUrl)
	failOnError(err, "Could not load DB configuration")

	dbConnection := pg.Connect(options)
	defer dbConnection.Close()

	logger := logrus.New()

	customerService := dependencyInjection(dbConnection, logger)
	customerController := controller.NewCustomerController(customerService, logger)

	err = http.ListenAndServe(fmt.Sprintf(":%s", servicePort), customerController.RestController())
	log.Fatal(err)
}

func dependencyInjection(dbConnection *pg.DB, logger *logrus.Logger) *service.CustomerService {
	customerRepository := repository.NewCustomerRepositoryImpl(dbConnection)
	customerService := service.NewCustomerService(customerRepository, logger)
	return customerService
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

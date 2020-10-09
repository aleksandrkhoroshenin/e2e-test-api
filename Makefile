.PHONY: install build run test migrate_init migrate migrate_down

API_NAME?=e2e-test-api
DB_PORT?=5480
DB_HOST?=stage.cobu.ru
DB_CONNECTION_URL?=postgres://$(API_NAME):$(API_NAME)@$(DB_HOST):$(DB_PORT)/$(API_NAME)?sslmode=disable

install:
	go mod tidy
	go mod download

build: install
	go build -o ./bin/$(API_NAME) ./cmd/gateway
	go build -o ./bin/migrations ./migrations/

run: build
	DB_CONNECTION_URL=$(DB_CONNECTION_URL) \
	./bin/$(API_NAME)

test:
	go test ./api/...

migrate_init: build
	DB_CONNECTION_URL=$(DB_CONNECTION_URL) \
	./bin/migrations init

migrate: migrate_init
	DB_CONNECTION_URL=$(DB_CONNECTION_URL) \
	./bin/migrations

migrate_down: build
	./bin/migrations down

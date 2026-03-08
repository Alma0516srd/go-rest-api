.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./ ..

.PHONY: run
run:
	go run ./cmd/apiserver -config-path configs/apiserver.toml

.PHONY: db-up
db-up:
	docker-compose up -d

.PHONY: db-down
db-down:
	docker-compose down

.PHONY: db-logs
db-logs:
	docker-compose logs -f db

.PHONY: migrate-up
migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5435/restapi_dev?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5435/restapi_dev?sslmode=disable" down

.DEFAULT_GOAL := build
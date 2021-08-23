POSTGRES_USER := user-development
DB_NAME := test-db
POSTGRES_PASSWORD := secret
POSTGRESQL_URL := postgres://localhost:5432/$(DB_NAME)?user=$(POSTGRES_USER)&password=$(POSTGRES_PASSWORD)&sslmode=disable
MIGRATIONS_PATH := internal/storage/pkg/postgres/sql

# Colorful output
color_off = \033[0m
color_cyan = \033[1;36m
color_green = \033[1;32m

define log_info
	@printf "$(color_cyan)$(1)$(color_off)\n"
endef
define log_success
	@printf "$(color_green)$(1)$(color_off)\n"
endef

.PHONY: postgresstart postgresstop createdb dropdb migrate_up migrate_down e2e proto test/ci test/all up swagger-validate swagger-serve install_deps lint tidy test

install_deps:
	go mod download
	go install github.com/golang/protobuf/protoc-gen-go@v1.4.3
	brew install golangci-lint
	brew install golang-migrate
	# go install github.com/go-swagger/go-swagger/cmd/swagger@v0.26.1
	# go install github.com/go-swagger/go-swagger/cmd/swagger@v0.26.1

up:
	go mod tidy
	go mod download
	docker-compose up --remove-orphans -d --build

lint:
	golangci-lint run ./...
	$(call log_success,Linting with golangci-lint succeeded!)

tidy:
	$(call log_info,Check that go.mod and go.sum don't contain any unnecessary dependency)
	go mod tidy -v
	git diff-index --quiet HEAD
	$(call log_success,Go mod check succeeded!)

test:
	$(call log_info,Run tests and check race conditions)
	# https://golang.org/doc/articles/race_detector.html
	go test -race -v ./... -cover
	$(call log_success,All tests succeeded)

create:
	docker create --name postgres13 -p 5432:5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=secret -e PGPASSWORD=$(POSTGRES_PASSWORD) postgres:13.3

start:
	docker start postgres13

stop:
	docker stop postgres13

remove:
	docker rm postgres13

createdb:
	docker exec -it postgres13 createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(DB_NAME)

dropdb:
	docker exec -it postgres13 dropdb --username=$(POSTGRES_USER) $(DB_NAME)

sleep:
	sleep 10

migrate_up:
	migrate -path $(MIGRATIONS_PATH) -database="$(POSTGRESQL_URL)" -verbose up

migrate_down:
	migrate -path $(MIGRATIONS_PATH) -database="$(POSTGRESQL_URL)" -verbose down 1

e2e:
	$(call log_info,Starting test environment:)
	go test ./... -v
	docker compose up -d
	TEST_INTEGRATION=TRUE go test ./... -v 
	docker compose down 

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/protos/account.proto

test/ci: test go-mod-tidy

workflow: create start sleep createdb sleep migrate_up migrate_down dropdb stop remove

test/all: test go-mod-tidy lint e2e

swagger-validate:
	swagger validate ./api/swagger.yml

swagger-serve:
	swagger serve ./api/swagger.yml

-include ./tests/e2e/e2e.mk

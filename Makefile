POSTGRES_HOST := localhost
POSTGRES_PORT := 5432
DB_NAME := e2e-db
POSTGRES_USER := user-development
POSTGRES_PASSWORD := secret
POSTGRESQL_URL := postgresql://$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB_NAME)?user=$(POSTGRES_USER)&password=$(POSTGRES_PASSWORD)&sslmode=disable
MIGRATIONS_PATH := migrations

.PHONY: e2e proto up swagger-validate swagger-serve install_deps test

install_deps:
	go mod download
	go install github.com/golang/protobuf/protoc-gen-go@v1.4.3
	brew install golang-migrate
	# go install github.com/go-swagger/go-swagger/cmd/swagger@v0.26.1
	# go install github.com/go-swagger/go-swagger/cmd/swagger@v0.26.1

up:
	go mod tidy
	go mod download
	docker-compose up --remove-orphans -d --build

down:
	docker compose -p goapi down

tidy:
	go mod tidy -v
	git diff-index --quiet HEAD

test:
	# https://golang.org/doc/articles/race_detector.html
	go test -race -v ./... -cover

proto:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative internal/proto/goapi.proto

build:
	go build ./...

e2e.up:
	docker compose --file docker-compose.e2e.yaml up \
					--remove-orphans \
					--build \
					--detach

e2e.down:
	docker compose --file docker-compose.e2e.yaml down

migrate.up:
	migrate -path="$(MIGRATIONS_PATH)" -database="$(POSTGRESQL_URL)" -verbose up

migrate.down:
	migrate -path="$(MIGRATIONS_PATH)" -database="$(POSTGRESQL_URL)" -verbose down

e2e.test:
	gotest -tags=e2e ./... -v

e2e: proto build e2e.up migrate.up e2e.test migrate.down e2e.down

swagger-validate:
	swagger validate ./api/swagger.yml

swagger-serve:
	swagger serve ./api/swagger.yml

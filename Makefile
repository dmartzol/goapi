POSTGRES_HOST := localhost
POSTGRES_PORT := 5432
DB_NAME := goapi-development
POSTGRES_USER := user-development
POSTGRES_PASSWORD := password
POSTGRESQL_URL := postgresql://$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB_NAME)?user=$(POSTGRES_USER)&password=$(POSTGRES_PASSWORD)&sslmode=disable
MIGRATIONS_PATH := migrations
MIGRATE_VERSION := v4.15.1

.PHONY: e2e proto up swagger-validate swagger-serve install_deps test

install_deps:
	go mod download
	go install github.com/golang/protobuf/protoc-gen-go@v1.4.3
	brew install golang-migrate

up:
	docker-compose up --remove-orphans -d --build

down:
	docker compose -p goapi down

tidy:
	go mod tidy -v
	go mod download

test:
	# https://golang.org/doc/articles/race_detector.html
	go test -race -v ./... -cover

proto:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative internal/proto/goapi.proto

e2e.up:
	docker compose --file docker-compose.e2e.yaml up \
					--remove-orphans \
					--build \
					--detach

e2e.down:
	docker compose --file docker-compose.e2e.yaml down


migrate.up:
	docker run -v $(shell pwd)/migrations:/migrations \
				--network host migrate/migrate:$(MIGRATE_VERSION) \
				-path="$(MIGRATIONS_PATH)" \
				-database="$(POSTGRESQL_URL)" \
				-verbose \
				up

migrate.down:
	docker run -v $(shell pwd)/migrations:/migrations \
				--network host migrate/migrate:$(MIGRATE_VERSION) \
				-path="$(MIGRATIONS_PATH)" \
				-database="$(POSTGRESQL_URL)" \
				-verbose \
				down 1

e2e.test:
	gotest -tags=e2e ./... -v

e2e: proto build e2e.up migrate.up e2e.test migrate.down e2e.down

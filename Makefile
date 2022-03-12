POSTGRES_HOST := localhost
POSTGRES_PORT := 5432
DB_NAME := goapi-development
POSTGRES_USER := user-development
POSTGRES_PASSWORD := password
POSTGRESQL_URL := postgresql://$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB_NAME)?user=$(POSTGRES_USER)&password=$(POSTGRES_PASSWORD)&sslmode=disable
MIGRATIONS_PATH := migrations
MIGRATE_VERSION := v4.15.1

.PHONY: install_deps up down tidy proto migrate.up migrate.down 

install_deps:
	go mod download
	go install github.com/golang/protobuf/protoc-gen-go@v1.4.3

up:
	docker-compose up --remove-orphans -d

down:
	docker compose -p goapi down

tidy:
	go mod tidy -v
	go mod download

proto:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative internal/proto/goapi.proto

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

-include e2e.mk

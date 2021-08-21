POSTGRESQL_URL := postgres://user-development@localhost:5432/development?sslmode=disable
MIGRATIONS_PATH := internal/storage/pkg/postgres/sql

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path $(MIGRATIONS_PATH) -database $(POSTGRESQL_URL) -verbose up

migratedown:
	migrate -path $(MIGRATIONS_PATH) -database $(POSTGRESQL_URL) -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown


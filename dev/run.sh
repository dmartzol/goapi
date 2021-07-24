#bin/bash

go mod tidy
go mod download

docker-compose up --remove-orphans -d --build

GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgresql://user-development@localhost/development?sslmode=disable goose -dir="./internal/storage/postgres/sql" up

# go run main.go

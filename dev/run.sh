#bin/bash

go mod tidy
go mod download

docker-compose up --remove-orphans -d --build

# go run main.go

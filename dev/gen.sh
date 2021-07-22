#!/bin/bash

# gRPC
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/protos/account.proto

# buf beta mod update
# buf generate

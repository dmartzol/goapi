package main

import (
	"log"
	"net"

	accountservice "github.com/dmartzol/api-template/cmd/services/accounts/service"
	pb "github.com/dmartzol/api-template/internal/protos"
	"google.golang.org/grpc"
)

func main() {
	devMode := true
	aS, err := accountservice.NewAccountsService(devMode)
	if err != nil {
		log.Fatalf("failed to create accounts service: %+v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountsServer(s, aS)
	lis, err := net.Listen("tcp", ":"+accountservice.Port)
	if err != nil {
		aS.Fatalf("failed to listen: %v", err)
	}
	aS.Infow("listening and serving", "host", "0.0.0.0", "port", accountservice.Port)
	if err := s.Serve(lis); err != nil {
		aS.Fatalf("failed to serve: %+v", err)
	}
}

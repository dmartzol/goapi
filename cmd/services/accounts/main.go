package main

import (
	"log"
	"net"

	"github.com/dmartzol/goapi/internal/proto"
	"github.com/dmartzol/goapi/internal/service"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

func main() {
	var config service.AccountsServiceConfig
	err := envconfig.Process("goapi", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	aS, err := service.NewAccountsService(config)
	if err != nil {
		log.Fatalf("failed to create accounts service: %+v", err)
	}

	s := grpc.NewServer()
	proto.RegisterAccountsServer(s, aS)
	lis, err := net.Listen("tcp", ":"+service.Port)
	if err != nil {
		aS.Fatalf("failed to listen: %v", err)
	}
	aS.Infow("listening and serving", "host", "0.0.0.0", "port", service.Port)
	if err := s.Serve(lis); err != nil {
		aS.Fatalf("failed to serve: %+v", err)
	}
}

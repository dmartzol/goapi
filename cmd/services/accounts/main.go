package main

import (
	"log"
	"net"

	accountservice "github.com/dmartzol/goapi/cmd/services/accounts/service"
	pb "github.com/dmartzol/goapi/internal/proto"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

func main() {
	var config accountservice.Config
	err := envconfig.Process("goapi", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	aS, err := accountservice.New(config)
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

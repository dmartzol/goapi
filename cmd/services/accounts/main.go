package main

import (
	"context"
	"log"
	"net"

	pb "github.com/dmartzol/api-template/internal/protos"
	"github.com/dmartzol/api-template/internal/storage/postgres"
	"google.golang.org/grpc"
)

type accountService struct {
	pb.UnimplementedAccountsServer
	*postgres.DB
}

func (s *accountService) Account(ctx context.Context, accID *pb.AccountRequest) (*pb.AccountMessage, error) {
	a, err := s.DB.AccountWithCredentials("", "")
	log.Printf("acc %+v", a)
	return nil, err
}

func main() {
	dbClient, err := postgres.NewDBClient()
	if err != nil {
		log.Fatalf("error initializing database: %+v", err)
	}
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	aS := accountService{
		DB: dbClient,
	}
	pb.RegisterAccountsServer(s, &aS)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}

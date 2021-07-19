package main

import (
	"context"
	"log"
	"net"

	"github.com/dmartzol/api-template/internal/mylogger"
	pb "github.com/dmartzol/api-template/internal/protos"
	"github.com/dmartzol/api-template/internal/storage/postgres"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	port = "50051"
)

func NewAccountsService(devMode bool) (*accountService, error) {
	dbClient, err := postgres.NewDBClient()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db client")
	}
	logger, err := mylogger.NewLogger(devMode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create logger")
	}
	a := accountService{
		DB:            dbClient,
		SugaredLogger: logger,
	}
	return &a, nil
}

type accountService struct {
	pb.UnimplementedAccountsServer
	*postgres.DB
	*zap.SugaredLogger
}

func (s *accountService) Account(ctx context.Context, accID *pb.AccountRequest) (*pb.AccountMessage, error) {
	a, err := s.DB.AccountWithCredentials("", "")
	log.Printf("acc %+v", a)
	return nil, err
}

func main() {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	aS, err := NewAccountsService(true)
	if err != nil {
		log.Fatalf("failed to create accounts service: %+v", err)
	}
	pb.RegisterAccountsServer(s, aS)
	aS.Infow("listening and serving", "host", "0.0.0.0", "port", port)
	if err := s.Serve(lis); err != nil {
		aS.Fatalf("failed to serve: %+v", err)
	}
}

package accountservice

import (
	"context"
	"log"

	"github.com/dmartzol/api-template/internal/mylogger"
	pb "github.com/dmartzol/api-template/internal/protos"
	"github.com/dmartzol/api-template/internal/storage/postgres"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	Port = "50051"
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

func (s *accountService) Run() error {
	return nil
}

func (s *accountService) Account(ctx context.Context, accountID *pb.AccountID) (*pb.AccountMessage, error) {
	a, err := s.DB.AccountWithCredentials("", "")
	log.Printf("acc %+v", a)
	return nil, err
}

func (s *accountService) AddAccount(ctx context.Context, addMessage *pb.AddAccountMessage) (*pb.AccountMessage, error) {
	a, err := s.DB.AccountWithCredentials("", "")
	log.Printf("acc %+v", a)
	return nil, err
}

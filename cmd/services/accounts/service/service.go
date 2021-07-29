package accountservice

import (
	"context"

	"github.com/dmartzol/api-template/internal/model"
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
	s.Errorw("not implemented", "function", "Account(ctx context.Context, accountID *pb.AccountID)")
	return nil, errors.Errorf("not implemented")
}

func (s *accountService) AddAccount(ctx context.Context, addMessage *pb.AddAccountMessage) (*pb.AccountMessage, error) {
	accountInsert := &model.Account{
		FirstName: addMessage.FirstName,
		LastName:  addMessage.LastName,
	}
	newAccount, err := s.DB.AddAccount(accountInsert)
	if err != nil {
		s.Errorw("failed to add acount", "error", err)
		return nil, errors.Wrap(err, "failed to add account")
	}
	accountMessage := pb.AccountMessage{
		FirstName: newAccount.FirstName,
		LastName:  newAccount.LastName,
	}
	return &accountMessage, nil
}

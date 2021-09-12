package accountservice

import (
	"context"

	"github.com/dmartzol/goapi/goapi"
	"github.com/dmartzol/goapi/internal/logger"
	pb "github.com/dmartzol/goapi/internal/proto"
	"github.com/dmartzol/goapi/internal/storage"
	"github.com/dmartzol/goapi/internal/storage/pkg/postgres"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	Port = "50051"
)

func New(dbname, dbusername, dbhostname string, humanReadableLogs bool) (*accountService, error) {
	dbClient, err := postgres.NewDBClient(dbname, dbusername, dbhostname)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db client")
	}
	storageClient := storage.New(dbClient)
	logger, err := logger.New(humanReadableLogs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create logger")
	}
	a := accountService{
		Storage:       storageClient,
		SugaredLogger: logger,
	}
	return &a, nil
}

type accountService struct {
	pb.UnimplementedAccountsServer
	*storage.Storage
	*zap.SugaredLogger
}

func (s *accountService) Run() error {
	return nil
}

func (s *accountService) Account(ctx context.Context, accountID *pb.AccountID) (*pb.Account, error) {
	s.Errorw("not implemented", "function", "Account(ctx context.Context, accountID *pb.AccountID)")
	return nil, errors.Errorf("not implemented")
}

func (s *accountService) AddAccount(ctx context.Context, addMessage *pb.AddAccountMessage) (*pb.Account, error) {
	accountInsert := &goapi.Account{
		FirstName: addMessage.FirstName,
		LastName:  addMessage.LastName,
	}
	newAccount, err := s.Storage.AddAccount(accountInsert)
	if err != nil {
		s.Errorw("failed to add acount", "error", err)
		return nil, errors.Wrap(err, "failed to add account")
	}
	accountMessage := pb.Account{
		FirstName: newAccount.FirstName,
		LastName:  newAccount.LastName,
	}
	return &accountMessage, nil
}

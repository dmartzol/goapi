package service

import (
	"context"

	"github.com/dmartzol/goapi/internal/logger"
	"github.com/dmartzol/goapi/internal/proto"
	"github.com/dmartzol/goapi/internal/storage"
	"github.com/dmartzol/goapi/internal/storage/pkg/postgres"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	Port = "50051"
)

type accountService struct {
	proto.UnimplementedAccountsServer
	*storage.Storage
	*zap.SugaredLogger
}

type AccountsServiceConfig struct {
	StructuredLogging bool
	DatabaseHostname  string
	DatabaseName      string
	DatabaseUsername  string
	DatabasePassword  string
	DatabasePort      int
}

func NewAccountsService(config AccountsServiceConfig) (*accountService, error) {
	logger, err := logger.New(config.StructuredLogging)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create logger")
	}

	logger.Infof("structured logging: %t", config.StructuredLogging)
	logger.Infof("database hostname: %s", config.DatabaseHostname)
	logger.Infof("database name: %s", config.DatabaseName)
	logger.Infof("database username: %s", config.DatabaseUsername)

	dbConfig := postgres.Config{
		Host:     config.DatabaseHostname,
		Name:     config.DatabaseName,
		User:     config.DatabaseUsername,
		Password: config.DatabasePassword,
		Port:     config.DatabasePort,
	}
	dbClient, err := postgres.NewWithWaitLoop(dbConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db client")
	}

	a := accountService{
		Storage:       storage.New(dbClient),
		SugaredLogger: logger,
	}

	return &a, nil
}

func (s *accountService) Run() error {
	return nil
}

func (s *accountService) Account(ctx context.Context, accountID *proto.AccountID) (*proto.Account, error) {
	s.Errorw("not implemented", "function", "Account(ctx context.Context, accountID *proto.AccountID)")
	return nil, errors.Errorf("not implemented")
}

func (s *accountService) AddAccount(ctx context.Context, addAccountMessage *proto.AddAccountMessage) (*proto.Account, error) {
	newAccount := addAccountMessage.ToCoreAccount()
	newAccount, err := s.Storage.AddAccount(newAccount)
	if err != nil {
		s.Errorw("failed to add acount", "error", err)
		return nil, errors.Wrap(err, "failed to add account")
	}
	pbAccount, err := proto.AccountProto(newAccount)
	if err != nil {
		s.Errorw("failed to convert acount", "error", err)
		return nil, errors.Wrap(err, "failed to convert account")
	}
	return pbAccount, nil
}

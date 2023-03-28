package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/dmartzol/goapi/internal/commands"
	"github.com/dmartzol/goapi/internal/logger"
	"github.com/dmartzol/goapi/internal/proto"
	"github.com/dmartzol/goapi/internal/storage"
	"github.com/dmartzol/goapi/internal/storage/pkg/postgres"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	app := cli.App{
		Name:   "accounts",
		Usage:  "run accounts service",
		Action: newAccountsServiceRun,
	}
	app.Flags = append(app.Flags, commands.CommonFlags...)
	app.Flags = append(app.Flags, commands.ServiceFlags...)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("error running app: %s", err)
	}
}

type accountService struct {
	proto.UnimplementedAccountsServer
	*storage.MacroStorage
	*zap.SugaredLogger
}

func newAccountsServiceRun(c *cli.Context) error {
	structuredLogging := c.Bool(commands.StructuredLoggingFlagName)
	port := c.String("port")
	databaseHostname := c.String(commands.DatabaseHostnameFlagName)
	databaseName := c.String("databaseName")
	databaseUser := c.String("databaseUser")
	databasePassword := c.String("databasePassword")
	databasePort := c.Int("databasePort")
	hostname := c.String("host")

	logger, err := logger.New(structuredLogging, "accounts_logger")
	if err != nil {
		return errors.Wrap(err, "failed to create logger")
	}

	logger.Infof("structured logging: %v", structuredLogging)
	logger.Infof("hostname: %s", hostname)
	logger.Infof("database hostname: %s", databaseHostname)
	logger.Infof("database name: %s", databaseName)
	logger.Infof("database username: %s", databaseUser)

	dbConfig := postgres.Config{
		Host:     databaseHostname,
		Name:     databaseName,
		User:     databaseUser,
		Password: databasePassword,
		Port:     databasePort,
	}
	dbClient, err := postgres.NewWithWaitLoop(dbConfig)
	if err != nil {
		return errors.Wrap(err, "failed to create db client")
	}

	a := accountService{
		MacroStorage:  storage.New(dbClient),
		SugaredLogger: logger,
	}

	s := grpc.NewServer()
	proto.RegisterAccountsServer(s, a)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		a.Fatalf("failed to listen: %v", err)
	}

	a.Infow("listening and serving", "host", hostname, "port", port)

	return s.Serve(lis)
}

func (s *accountService) Account(ctx context.Context, accountID *proto.AccountID) (*proto.Account, error) {
	s.Errorw("not implemented", "function", "Account(ctx context.Context, accountID *proto.AccountID)")
	return nil, errors.Errorf("not implemented")
}

func (s accountService) AddAccount(ctx context.Context, addAccountMessage *proto.AddAccountMessage) (*proto.Account, error) {
	newAccount := addAccountMessage.ToCoreAccount()
	newAccount, err := s.MacroStorage.AddAccount(newAccount)
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

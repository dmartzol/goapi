package gatewayservice

import (
	"log"

	accountservice "github.com/dmartzol/api-template/cmd/services/accounts/service"
	"github.com/dmartzol/api-template/internal/handler"
	"github.com/dmartzol/api-template/internal/mylogger"
	pb "github.com/dmartzol/api-template/internal/protos"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	accountsHost = "accounts"
)

type gatewayService struct {
	*handler.Handler
}

func NewGatewayService(devMode bool) (*gatewayService, error) {
	logger, err := mylogger.NewLogger(devMode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create logger")
	}
	logger.Info("creating database client")
	logger.Info("creating grcp connection")
	accountsAddres := accountsHost + ":" + accountservice.Port
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(accountsAddres, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create connection")
	}
	// defer conn.Close()
	logger.Info("creating accounts service client")
	accountsClient := pb.NewAccountsClient(conn)
	apiHandler, err := handler.NewHandler(accountsClient, logger)
	if err != nil {
		log.Panicf("error creating handler: %v", err)
	}
	logger.Info("initializing routes")
	apiHandler.InitializeRoutes()
	s := gatewayService{
		Handler: apiHandler,
	}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	return &s, nil
}

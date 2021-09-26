package gatewayservice

import (
	"log"

	accountservice "github.com/dmartzol/goapi/cmd/services/accounts/service"
	"github.com/dmartzol/goapi/internal/handler"
	"github.com/dmartzol/goapi/internal/logger"
	pb "github.com/dmartzol/goapi/internal/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	accountsHost = "accounts"
)

type gatewayService struct {
	*handler.Handler
}

func New(structuredLogging, verbose bool) (*gatewayService, error) {
	logger, err := logger.New(structuredLogging)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create logger")
	}
	accountsAddres := accountsHost + ":" + accountservice.Port
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(accountsAddres, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create connection")
	}
	// defer conn.Close()
	accountsClient := pb.NewAccountsClient(conn)
	handler, err := handler.New(accountsClient, logger, verbose)
	if err != nil {
		log.Panicf("error creating handler: %v", err)
	}
	handler.InitializeRoutes()
	s := gatewayService{
		Handler: handler,
	}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	return &s, nil
}

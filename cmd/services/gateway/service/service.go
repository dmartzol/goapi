package gatewayservice

import (
	"log"

	"github.com/dmartzol/api-template/internal/handler"
	"github.com/dmartzol/api-template/internal/mylogger"
	pb "github.com/dmartzol/api-template/internal/protos"
	"github.com/dmartzol/api-template/internal/storage/postgres"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	accountsHost = "accounts"
	accountsPort = "50051"
)

type GatewayService struct {
	*handler.Handler
}

func NewGatewayService(devMode bool) (*GatewayService, error) {
	logger, err := mylogger.NewLogger(devMode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create logger")
	}
	dbClient, err := postgres.NewDBClient()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db client")
	}
	accountsAddres := accountsHost + ":" + accountsPort
	conn, err := grpc.Dial(accountsAddres, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create connection")
	}
	defer conn.Close()
	accountsClient := pb.NewAccountsClient(conn)
	apiHandler, err := handler.NewHandler(dbClient, accountsClient, logger)
	if err != nil {
		log.Panicf("error creating handler: %v", err)
	}
	apiHandler.InitializeRoutes()
	s := GatewayService{
		Handler: apiHandler,
	}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	return &s, nil
}

func (g *GatewayService) ListenAndServe() error {
	return g.Run("0.0.0.0:1100")
}

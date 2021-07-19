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
	address = "localhost:50051"
)

type GatewayService struct {
	*handler.Handler
	AccountsClient *pb.AccountsClient
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
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create connection")
	}
	defer conn.Close()
	apiHandler, err := handler.NewHandler(dbClient, logger)
	if err != nil {
		log.Panicf("error creating handler: %v", err)
	}
	apiHandler.InitializeRoutes()

	accountsClient := pb.NewAccountsClient(conn)
	s := GatewayService{
		AccountsClient: &accountsClient,
	}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	return &s, nil
}

func (g *GatewayService) ListenAndServe() error {
	return g.Run("0.0.0.0:1100")
}

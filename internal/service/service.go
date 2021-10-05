package service

import (
	"log"
	"net/http"
	"strings"

	accountservice "github.com/dmartzol/goapi/cmd/services/accounts/service"
	"github.com/dmartzol/goapi/internal/handler"
	"github.com/dmartzol/goapi/internal/logger"
	pb "github.com/dmartzol/goapi/internal/proto"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

const (
	accountsHost = "accounts"
)

type Config struct {
	StructuredLogging bool
	Verbose           bool
}

func NewGatewayServiceRun(c *cli.Context) error {
	structuredLogging := c.Bool("structuredLogging")
	verbose := c.Bool("verbose")
	hostname := c.String("hostname")
	port := c.String("port")

	logger, err := logger.New(structuredLogging)
	if err != nil {
		return errors.Wrap(err, "failed to create logger")
	}
	accountsAddres := accountsHost + ":" + accountservice.Port
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(accountsAddres, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to create connection")
	}
	defer conn.Close()

	accountsClient := pb.NewAccountsClient(conn)
	handler, err := handler.New(accountsClient, logger, verbose)
	if err != nil {
		log.Panicf("error creating handler: %v", err)
	}
	handler.InitializeRoutes()

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	handler.Infof("listening and serving on %s", hostname)
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})
	address := strings.Join([]string{hostname, port}, ":")
	return http.ListenAndServe(address, cors.Handler(handler.Router))
}

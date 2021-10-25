package service

import (
	"log"
	"net/http"

	"github.com/dmartzol/goapi/internal/handler"
	"github.com/dmartzol/goapi/internal/logger"
	"github.com/dmartzol/goapi/internal/proto"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func NewGatewayServiceRun(c *cli.Context) error {
	structuredLogging := c.Bool("structuredLogging")
	verbose := c.Bool("verbose")
	hostname := c.String("host")
	port := c.String("port")
	accountsServiceHost := c.String("accountsServiceHostname")
	accountsServicePort := c.String("accountsServicePort")

	logger, err := logger.New(structuredLogging)
	if err != nil {
		return errors.Wrap(err, "failed to create logger")
	}

	accountsAddres := accountsServiceHost + ":" + accountsServicePort
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(accountsAddres, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to create connection")
	}
	defer conn.Close()
	accountsClient := proto.NewAccountsClient(conn)

	handler, err := handler.New(accountsClient, logger, verbose)
	if err != nil {
		log.Panicf("error creating handler: %v", err)
	}
	handler.InitializeRoutes()

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	handler.Infof("listening and serving on %s:%s", hostname, port)
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})

	address := hostname + ":" + port
	return http.ListenAndServe(address, cors.Handler(handler.Router))
}

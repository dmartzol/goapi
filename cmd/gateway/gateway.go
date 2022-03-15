package main

import (
	"log"

	"github.com/dmartzol/goapi/internal/commands"
	"github.com/dmartzol/goapi/internal/handler"
	"github.com/dmartzol/goapi/internal/logger"
	"github.com/dmartzol/goapi/internal/proto"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func newGatewayServiceRun(c *cli.Context) error {
	hostname := c.String(commands.HostnameFlagName)
	port := c.String(commands.PortFlagName)
	structuredLogging := c.Bool(commands.StructuredLoggingFlagName)
	rawRequestLogging := c.Bool(commands.RawRequestsLoggingFlagName)
	accountsServiceHost := c.String("accountsServiceHostname")
	accountsServicePort := c.String("accountsServicePort")

	logger, err := logger.New(structuredLogging, "gateway_logger")
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

	handler, err := handler.New(accountsClient, logger, rawRequestLogging)
	if err != nil {
		log.Panicf("error creating handler: %v", err)
	}
	handler.InitializeRoutes()

	// Port details: https://www.jaegertracing.io/docs/getting-started/
	//je, err := jaeger.NewExporter(jaeger.Options{
	//AgentEndpoint:     "localhost:6831",
	//CollectorEndpoint: "http://localhost:14268/api/traces",
	//ServiceName:       "my_service",
	//})
	//if err != nil {
	//log.Fatalf("Failed to create the Jaeger exporter: %v", err)
	//}

	// And now finally register it as a Trace Exporter
	//trace.RegisterExporter(je)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	//cors := cors.New(cors.Options{
	//AllowedOrigins:   []string{"http://localhost:3000"},
	//AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	//AllowCredentials: true,
	//// Enable Debugging for testing, consider disabling in production
	//// Debug: true,
	//})

	handler.Infof("listening and serving on %s:%s", hostname, port)
	address := hostname + ":" + port
	return handler.Router.Run(address)
}

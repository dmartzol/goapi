package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name: "gateway",
		Flags: []cli.Flag{
			&cli.BoolTFlag{
				Name:   "structuredLogin",
				EnvVar: "STRUCTURED_LOGIN",
			},
			&cli.StringFlag{
				Name:   "host",
				EnvVar: "HOST",
				Value:  "0.0.0.0",
			},
			&cli.StringFlag{
				Name:   "port",
				EnvVar: "PORT",
				Value:  "1100",
			},
			&cli.StringFlag{
				Name:   "accountsServiceHostname",
				EnvVar: "ACCOUNTS_SERVICE_HOSTNAME",
				Value:  "accounts",
			},
			&cli.StringFlag{
				Name:   "accountsServicePort",
				EnvVar: "ACCOUNTS_SERVICE_PORT",
				Value:  "50051",
			},
		},
		Action: newGatewayServiceRun,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("error running app: %s", err)
	}
}

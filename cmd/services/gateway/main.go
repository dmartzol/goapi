package main

import (
	"os"

	"github.com/dmartzol/goapi/internal/service"
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
				Name:   "hostname",
				EnvVar: "HOSTNAME",
				Value:  "localhost",
			},
			&cli.StringFlag{
				Name:   "port",
				EnvVar: "PORT",
				Value:  "1100",
			},
		},
		Action: service.NewGatewayServiceRun,
	}
	app.Run(os.Args)
}

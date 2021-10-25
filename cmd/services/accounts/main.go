package main

import (
	"log"
	"os"

	"github.com/dmartzol/goapi/internal/service"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name: "accounts",
		Flags: []cli.Flag{
			&cli.BoolTFlag{
				Name:   "structuredLogin",
				EnvVar: "STRUCTURED_LOGIN",
			},
			&cli.StringFlag{
				Name:   "host",
				EnvVar: "HOST",
				Value:  "localhost",
			},
			&cli.StringFlag{
				Name:   "port",
				EnvVar: "PORT",
				Value:  "50051",
			},
			&cli.StringFlag{
				Name:   "databaseHostname",
				EnvVar: "PGHOST",
				Value:  "database",
			},
			&cli.StringFlag{
				Name:   "databasePort",
				EnvVar: "PGPORT",
				Value:  "5432",
			},
			&cli.StringFlag{
				Name:   "databaseUser",
				EnvVar: "PGUSER",
				Value:  "user-development",
			},
			&cli.StringFlag{
				Name:   "databaseName",
				EnvVar: "PGNAME",
				Value:  "development",
			},
			&cli.StringFlag{
				Name:   "databasePassword",
				EnvVar: "PGPASSWORD",
				Value:  "development",
			},
		},
		Action: service.NewAccountsService,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("error running app: %s", err)
	}

}

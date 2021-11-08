package commands

import "github.com/urfave/cli"

var (
	StructuredLoggingFlagName  = "structuredLogging"
	RawRequestsLoggingFlagName = "rawRequestsLogging"
	DatabaseHostnameFlagName   = "databaseHostname"

	CommonFlags = []cli.Flag{
		&cli.BoolTFlag{
			Name:   StructuredLoggingFlagName,
			Usage:  "enable structured logging",
			EnvVar: "STRUCTURED_LOGGING",
		},
		&cli.StringFlag{
			Name:   DatabaseHostnameFlagName,
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
	}

	ServiceFlags = []cli.Flag{
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
	}
)

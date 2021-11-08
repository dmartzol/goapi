package main

import (
	"log"
	"os"

	"github.com/dmartzol/goapi/internal/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.App{
		Name:   "accounts",
		Usage:  "run accounts service",
		Action: newAccountsServiceRun,
	}
	app.Flags = append(app.Flags, commands.CommonFlags...)
	app.Flags = append(app.Flags, commands.ServiceFlags...)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("error running app: %s", err)
	}
}

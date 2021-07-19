package main

import (
	"log"

	accountservice "github.com/dmartzol/api-template/cmd/services/accounts/service"
	"github.com/dmartzol/api-template/internal/handler"
	"github.com/dmartzol/api-template/internal/storage/postgres"
)

func main() {
	dbClient, err := postgres.NewDBClient()
	if err != nil {
		log.Fatalf("error initializing database: %+v", err)
	}

	apiHandler, err := handler.NewHandler(dbClient, true)
	if err != nil {
		log.Panicf("error creating handler: %v", err)
	}
	apiHandler.InitializeRoutes()
	apiHandler.Infof("accounts service running on port %s", accountservice.Port)
	apiHandler.Run("0.0.0.0:1100")
}

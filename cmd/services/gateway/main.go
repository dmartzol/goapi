package main

import (
	"log"

	"github.com/dmartzol/api-template/internal/handler"
	"github.com/dmartzol/api-template/internal/storage/postgres"
)

func main() {
	dbClient, err := postgres.NewDBClient()
	if err != nil {
		log.Fatalf("error initializing database: %+v", err)
	}

	apiHandler := handler.NewHandler(dbClient, true)
	apiHandler.InitializeRoutes()
	apiHandler.Run("0.0.0.0:1100")
}

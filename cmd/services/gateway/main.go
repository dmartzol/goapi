package main

import (
	"log"

	"github.com/dmartzol/api-template/internal/handler"
	"github.com/dmartzol/api-template/internal/storage/postgres"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("error initializing logger: %+v", err)
	}
	defer logger.Sync()
	suggaredLogger := logger.Sugar()

	dbClient, err := postgres.NewDBClient()
	if err != nil {
		suggaredLogger.Fatalf("error initializing database: %+v", err)
	}

	apiHandler := handler.NewHandler(mux.NewRouter(), dbClient, suggaredLogger)
	apiHandler.InitializeRoutes()
	apiHandler.Run("0.0.0.0:1100")
}

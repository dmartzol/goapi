package main

import (
	"log"

	"github.com/dmartzol/api-template/internal/handler"
	"github.com/dmartzol/api-template/internal/storage/postgres"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	Ldate         = 1 << iota                  // the date in the local time zone: 2009/01/23
	Ltime                                      // the time in the local time zone: 01:23:23
	Lmicroseconds                              // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                                  // full file name and line number: /a/b/c/d.go:23
	Lshortfile                                 // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                                       // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                                 // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime | Lshortfile // initial values for the standard logger
)

func main() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("error initializing logger: %+v", err)
	}
	defer l.Sync()
	logger := l.Sugar()

	log.SetFlags(LstdFlags)
	dbClient, err := postgres.NewDBClient()
	if err != nil {
		log.Fatalf("error initializing database: %+v", err)
	}

	apiHandler := handler.NewHandler(mux.NewRouter(), dbClient, logger)
	apiHandler.InitializeRoutes()
	apiHandler.Run("0.0.0.0:1100")
}

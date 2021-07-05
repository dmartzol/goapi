package main

import (
	"log"
	"net/http"

	"github.com/dmartzol/api-template/internal/handler"
	"github.com/dmartzol/api-template/internal/storage/postgres"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	log.SetFlags(LstdFlags)
	db, err := postgres.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	handler, err := handler.NewHandler(db)
	if err != nil {
		log.Fatalf("error starting api: %+v", err)
	}

	r := mux.NewRouter()
	r = r.PathPrefix("/v1").Subrouter()

	r.HandleFunc("/version", handler.Version).Methods("GET")

	// sessions
	// see: https://stackoverflow.com/questions/7140074/restfully-design-login-or-register-resources
	r.HandleFunc("/sessions", handler.CreateSession).Methods("POST")
	r.HandleFunc("/sessions", handler.GetSession).Methods("GET")
	r.HandleFunc("/sessions", handler.ExpireSession).Methods("DELETE")

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		handler.AuthMiddleware,
	)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})

	log.Print("listening and serving")
	log.Fatal(http.ListenAndServe("localhost:3001", c.Handler(r)))
}

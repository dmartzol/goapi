package main

import (
	"log"

	gatewayservice "github.com/dmartzol/goapi/cmd/services/gateway/service"
)

func main() {
	structuredLogging := false
	verbose := true

	g, err := gatewayservice.New(structuredLogging, verbose)
	if err != nil {
		log.Fatalf("failed to create gateway: %v", err)
	}
	log.Fatal(g.Run("0.0.0.0:1100"))
}

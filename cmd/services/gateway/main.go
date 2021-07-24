package main

import (
	"log"

	gatewayservice "github.com/dmartzol/api-template/cmd/services/gateway/service"
)

func main() {
	devMode := true
	g, err := gatewayservice.NewGatewayService(devMode)
	if err != nil {
		log.Fatalf("failed to create gateway: %v", err)
	}
	log.Fatal(g.Run("0.0.0.0:1100"))
}

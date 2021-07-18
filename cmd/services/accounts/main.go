package main

import (
	"log"

	"github.com/dmartzol/api-template/internal/handler"
	"github.com/dmartzol/api-template/internal/proto/protoAccount"
	"github.com/dmartzol/api-template/internal/storage/postgres"
)

// server is used to implement helloworld.GreeterServer.
// type accountsService struct {
// 	pb.UnimplementedGreeterServer
// }

type accountService struct {
	protoAccount.UnimplementedAccountsServer
	*postgres.DB
}

// func main() {
// 	log.Fatal("")
// 	a := model.Account{}
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterGreeterServer(s, &accountsService{})

// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

func main() {
	dbClient, err := postgres.NewDBClient()
	if err != nil {
		log.Fatalf("error initializing database: %+v", err)
	}

	apiHandler := handler.NewHandler(dbClient, true)
	apiHandler.InitializeRoutes()
	apiHandler.Run("0.0.0.0:1100")
}

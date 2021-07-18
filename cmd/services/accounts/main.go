package main

import (
	pb "github.com/dmartzol/api-template/internal/protos"
	"github.com/dmartzol/api-template/internal/storage/postgres"
)

// server is used to implement helloworld.GreeterServer.
// type accountsService struct {
// 	pb.UnimplementedGreeterServer
// }

type accountService struct {
	pb.UnimplementedAccountsServer
	*postgres.DB
}

// func (s *accountService) Account(ctx context.Context, accID *protoTemplate.AccountRequest) (*protoTemplate.AccountMessage, error) {
// 	a, err := s.DB.AccountWithCredentials("", "")
// 	return a, err
// }

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
	// a := template.ThisConst
	// log.Printf("my constant %s", a)

	// dbClient, err := postgres.NewDBClient()
	// if err != nil {
	// 	log.Fatalf("error initializing database: %+v", err)
	// }
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
}

package main

import (
	"flag"
	"log"
	"net"

	accountservice "github.com/dmartzol/api-template/cmd/services/accounts/service"
	pb "github.com/dmartzol/api-template/internal/protos"
	"google.golang.org/grpc"
)

func main() {
	var (
		structuredLogging = flag.Bool("d", false, "")
		dbhostname        = flag.String("dbhostname", "database", "")
		dbusername        = flag.String("dbusername", "user-development", "")
		dbname            = flag.String("dbname", "database", "")
	)
	flag.Parse()
	aS, err := accountservice.NewAccountsService(
		*dbname,
		*dbusername,
		*dbhostname,
		*structuredLogging,
	)
	if err != nil {
		log.Fatalf("failed to create accounts service: %+v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountsServer(s, aS)
	lis, err := net.Listen("tcp", ":"+accountservice.Port)
	if err != nil {
		aS.Fatalf("failed to listen: %v", err)
	}
	aS.Infow("listening and serving", "host", "0.0.0.0", "port", accountservice.Port)
	if err := s.Serve(lis); err != nil {
		aS.Fatalf("failed to serve: %+v", err)
	}
}

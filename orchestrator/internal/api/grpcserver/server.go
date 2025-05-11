package grpcserver

import (
	"database/sql"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/bootstrap"
	pb "github.com/SteeperMold/Calculator-go/orchestrator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunGRPCServer(db *sql.DB, config *bootstrap.Config) {
	lis, err := net.Listen("tcp", ":"+config.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	orchestratorService := NewOrchestratorService(db, config)
	pb.RegisterOrchestratorServer(grpcServer, orchestratorService)

	log.Printf("grpc listening on port %s", config.GRPCPort)
	log.Fatal(grpcServer.Serve(lis))
}

package application

import (
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/bootstrap"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	pb "github.com/SteeperMold/Calculator-go/orchestrator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

type Application struct {
	Config      *bootstrap.Config
	expressions []domain.Expression
	mu          sync.Mutex
	pb.UnimplementedOrchestratorServer
}

func New() *Application {
	return &Application{
		Config: bootstrap.NewConfigFromEnv(),
	}
}

func (a *Application) RunHTTPServer() {
	log.Printf("http running on port %s", a.Config.HTTPPort)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", a.PostExpressionHandler)
	mux.HandleFunc("/api/v1/expressions/", a.GetExpressionHandler)
	mux.HandleFunc("/api/v1/expressions", a.ExpressionListHandler)

	log.Fatal(http.ListenAndServe(":"+a.Config.HTTPPort, corsMiddleware(mux)))
}

func (a *Application) RunGRPCServer() {
	lis, err := net.Listen("tcp", ":"+a.Config.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrchestratorServer(grpcServer, a)

	log.Printf("grpc listening on port %s", a.Config.GRPCPort)
	log.Fatal(grpcServer.Serve(lis))
}

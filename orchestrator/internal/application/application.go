package application

import (
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"log"
	"net/http"
	"sync"
)

type Application struct {
	Config      *Config
	expressions []domain.Expression
	mu          sync.Mutex
}

func New() *Application {
	return &Application{
		Config: NewConfigFromEnv(),
	}
}

func (a *Application) RunServer() {
	log.Printf("Running on port %s", a.Config.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", a.PostExpressionHandler)
	mux.HandleFunc("/api/v1/expressions/", a.GetExpressionHandler)
	mux.HandleFunc("/api/v1/expressions", a.ExpressionListHandler)
	mux.HandleFunc("/internal/task", a.TaskHandler)

	log.Fatal(http.ListenAndServe(":"+a.Config.Port, corsMiddleware(mux)))
}

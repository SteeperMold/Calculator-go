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
	log.Printf("Running on port %s", a.Config.Address)
	http.HandleFunc("/api/v1/calculate", a.PostExpressionHandler)
	http.HandleFunc("/api/v1/expressions/", a.GetExpressionHandler)
	http.HandleFunc("/api/v1/expressions", a.ExpressionListHandler)
	http.HandleFunc("/internal/task", a.TaskHandler)
	log.Fatal(http.ListenAndServe(":"+a.Config.Address, nil))
}

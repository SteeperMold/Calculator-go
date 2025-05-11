package main

import (
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/api/grpcserver"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/api/httpserver"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/bootstrap"
)

func main() {
	app := bootstrap.NewApp()
	defer app.CloseDatabase()

	go func() {
		grpcserver.RunGRPCServer(app.DB, app.Config)
	}()

	httpserver.RunHTTPServer(app.DB, app.Config)
}

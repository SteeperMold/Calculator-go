package main

import "github.com/SteeperMold/Calculator-go/orchestrator/internal/application"

func main() {
	app := application.New()

	go func() {
		app.RunHTTPServer()
	}()

	app.RunGRPCServer()
}

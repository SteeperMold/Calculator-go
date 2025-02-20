package main

import "github.com/SteeperMold/Calculator-go/orchestrator/internal/application"

func main() {
	app := application.New()
	app.RunServer()
}

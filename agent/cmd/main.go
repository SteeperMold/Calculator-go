package main

import "github.com/SteeperMold/Calculator-go/agent/internal/application"

func main() {
	app := application.New()
	app.RunDaemon()
}

package httpserver

import (
	"database/sql"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/bootstrap"
	"log"
	"net/http"
)

func RunHTTPServer(db *sql.DB, config *bootstrap.Config) {
	log.Printf("http running on port %s", config.HTTPPort)

	mux := http.NewServeMux()
	jwtAuthMiddleware := jwtAuthMiddlewareBuilder(config.AccessTokenSecret)

	eh := NewExpressionHandler(db, config.ContextTimeout)
	mux.Handle("/api/v1/calculate", jwtAuthMiddleware(http.HandlerFunc(eh.CreateExpression)))
	mux.HandleFunc("/api/v1/expressions/", eh.GetExpression)
	mux.HandleFunc("/api/v1/expressions", eh.GetExpressionsList)

	uh := NewUsersHandler(db, config)
	mux.Handle("/api/v1/profile", jwtAuthMiddleware(http.HandlerFunc(uh.Profile)))
	mux.HandleFunc("/api/v1/register", uh.Signup)
	mux.HandleFunc("/api/v1/login", uh.Login)
	mux.HandleFunc("/api/v1/refresh", uh.RefreshToken)

	log.Fatal(http.ListenAndServe(":"+config.HTTPPort, corsMiddleware(mux)))
}

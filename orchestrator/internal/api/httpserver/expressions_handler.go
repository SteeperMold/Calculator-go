package httpserver

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/service"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ExpressionHandler struct {
	service        *service.ExpressionsService
	contextTimeout time.Duration
}

func NewExpressionHandler(db *sql.DB, contextTimeout time.Duration) *ExpressionHandler {
	return &ExpressionHandler{
		service:        service.NewExpressionsService(db, contextTimeout),
		contextTimeout: contextTimeout,
	}
}

func (eh *ExpressionHandler) CreateExpression(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.NewExpressionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}

	newExprId, err := eh.service.CreateExpression(req.Expression)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidExpression) {
			http.Error(w, "invalid expression", http.StatusUnprocessableEntity)
			return
		}

		log.Printf("failed to create expression: %v\n", err)
		http.Error(w, "failed to create expression", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&domain.NewExpressionResponse{ID: *newExprId})
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (eh *ExpressionHandler) GetExpressionsList(w http.ResponseWriter, r *http.Request) {
	expressions, err := eh.service.GetExpressionsList()
	if err != nil {
		log.Printf("failed to get expressions list: %v\n", err)
		http.Error(w, "failed to get expressions list", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&domain.ExpressionsListResponse{Expressions: expressions})
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		return
	}
}

func (eh *ExpressionHandler) GetExpression(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")

	idStr := urlParts[4]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	selectedExpr, err := eh.service.GetExpressionById(int64(id))
	if err != nil {
		if errors.Is(err, domain.ErrExpressionDoesntExist) {
			http.Error(w, "expression doesn't exist", http.StatusNotFound)
			return
		}

		log.Printf("failed to get expression: %v\n", err)
		http.Error(w, "failed to get expression", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(selectedExpr)
	if err != nil {
		http.Error(w, "failed to encode expression", http.StatusInternalServerError)
		return
	}
}

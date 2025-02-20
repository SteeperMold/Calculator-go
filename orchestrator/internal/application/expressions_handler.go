package application

import (
	"encoding/json"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"github.com/SteeperMold/Calculator-go/orchestrator/pkg/calculation"
	"net/http"
	"strconv"
	"strings"
)

type NewExpressionRequest struct {
	Expression string `json:"expression"`
}

type NewExpressionResponse struct {
	ID int `json:"id"`
}

func (a *Application) PostExpressionHandler(w http.ResponseWriter, r *http.Request) {
	var req NewExpressionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	ast, err := calculation.BuildAST(req.Expression)
	if err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}

	newExpr := domain.Expression{
		ID:     len(a.expressions),
		Status: domain.StatusInProgress,
		Result: 0,
		AST:    ast,
	}

	a.expressions = append(a.expressions, newExpr)

	err = json.NewEncoder(w).Encode(&NewExpressionResponse{ID: newExpr.ID})
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type ExpressionsListResponse struct {
	Expressions []domain.Expression `json:"expressions"`
}

func (a *Application) ExpressionListHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&ExpressionsListResponse{Expressions: a.expressions})
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		return
	}
}

func (a *Application) GetExpressionHandler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")

	idStr := urlParts[4]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	var selectedExpr *domain.Expression

	for i := range a.expressions {
		if a.expressions[i].ID == id {
			selectedExpr = &a.expressions[i]
			break
		}
	}

	if selectedExpr == nil {
		http.Error(w, "expression not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(selectedExpr)
	if err != nil {
		http.Error(w, "failed to encode expression", http.StatusInternalServerError)
		return
	}
}

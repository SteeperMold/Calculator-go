package application

import (
	"encoding/json"
	"fmt"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"github.com/SteeperMold/Calculator-go/orchestrator/pkg/calculation"
	"net/http"
	"strconv"
)

type Task struct {
	ExpressionID  int     `json:"expression_id"`
	NodeID        int     `json:"node_id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime int     `json:"operation_time"`
}

type GetTaskResponse struct {
	Task Task `json:"task"`
}

func (a *Application) handleGetTask(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()

	var selectedExpr *domain.Expression

	for i := range a.expressions {
		if a.expressions[i].Status == domain.StatusInProgress {
			selectedExpr = &a.expressions[i]
			break
		}
	}

	if selectedExpr == nil {
		http.Error(w, "no tasks available", http.StatusNotFound)
		return
	}

	deepestNode, exists := calculation.GetDeepestInProgressOperation(selectedExpr.AST)
	if !exists {
		http.Error(w, "no tasks available", http.StatusNotFound)
		return
	}

	arg1, err := strconv.ParseFloat(deepestNode.Left.Value, 64)
	if err != nil {
		http.Error(w, "failed to parse float", http.StatusInternalServerError)
		return
	}
	arg2, err := strconv.ParseFloat(deepestNode.Right.Value, 64)
	if err != nil {
		http.Error(w, "failed to parse float", http.StatusInternalServerError)
		return
	}

	operation := deepestNode.Value
	var operationTime int

	switch operation {
	case "+":
		operationTime = a.Config.TimeAdditionMs
	case "-":
		operationTime = a.Config.TimeSubtractionMs
	case "*":
		operationTime = a.Config.TimeMultiplicationMs
	case "/":
		operationTime = a.Config.TimeDivisionMs
	default:
		operationTime = 0
	}

	response := &GetTaskResponse{Task: Task{
		ExpressionID:  selectedExpr.ID,
		NodeID:        deepestNode.ID,
		Arg1:          arg1,
		Arg2:          arg2,
		Operation:     deepestNode.Value,
		OperationTime: operationTime,
	}}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
	deepestNode.Status = domain.StatusGivenToAgent
}

type PostTaskRequest struct {
	ExpressionID int     `json:"expression_id"`
	NodeID       int     `json:"node_id"`
	Result       float64 `json:"result"`
}

func (a *Application) handlePostTask(w http.ResponseWriter, r *http.Request) {
	var req PostTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	var selectedExpr *domain.Expression

	for i := range a.expressions {
		if a.expressions[i].ID == req.ExpressionID {
			selectedExpr = &a.expressions[i]
			break
		}
	}

	if selectedExpr == nil {
		http.Error(w, "expression not found", http.StatusNotFound)
		return
	}

	if selectedExpr.Status != domain.StatusInProgress {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	calculatedNode, exists := calculation.GetNodeByID(selectedExpr.AST, req.NodeID)
	if !exists {
		http.Error(w, "node not found", http.StatusNotFound)
		return
	}

	if calculatedNode.Status != domain.StatusGivenToAgent {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	calculatedNode.Value = fmt.Sprintf("%.2f", req.Result)
	calculatedNode.Left = nil
	calculatedNode.Right = nil
	calculatedNode.Status = domain.StatusFinished

	if calculatedNode == selectedExpr.AST {
		selectedExpr.Status = domain.StatusFinished
		selectedExpr.Result = req.Result
		selectedExpr.AST = nil
	}
}

func (a *Application) TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.handleGetTask(w, r)
	case http.MethodPost:
		a.handlePostTask(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

package application

import (
	"context"
	"fmt"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"github.com/SteeperMold/Calculator-go/orchestrator/pkg/calculation"
	pb "github.com/SteeperMold/Calculator-go/orchestrator/proto"
	"strconv"
)

func (a *Application) FetchTask(ctx context.Context, _ *pb.Empty) (*pb.GetTaskResponse, error) {
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
		return nil, fmt.Errorf("no tasks available")
	}

	deepestNode, exists := calculation.GetDeepestInProgressOperation(selectedExpr.AST)
	if !exists {
		return nil, fmt.Errorf("no tasks available")
	}

	arg1, err := strconv.ParseFloat(deepestNode.Left.Value, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse float")
	}
	arg2, err := strconv.ParseFloat(deepestNode.Right.Value, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse float")
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

	deepestNode.Status = domain.StatusGivenToAgent

	return &pb.GetTaskResponse{
		ExpressionId:  int32(selectedExpr.ID),
		NodeId:        int32(deepestNode.ID),
		Arg1:          arg1,
		Arg2:          arg2,
		Operation:     deepestNode.Value,
		OperationTime: int32(operationTime),
	}, nil
}

func (a *Application) SendResult(ctx context.Context, tr *pb.PostTaskResult) (*pb.Empty, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	var selectedExpr *domain.Expression

	for i := range a.expressions {
		if a.expressions[i].ID == int(tr.ExpressionId) {
			selectedExpr = &a.expressions[i]
			break
		}
	}

	if selectedExpr == nil {
		return nil, fmt.Errorf("expression not found")
	}

	if selectedExpr.Status != domain.StatusInProgress {
		return nil, fmt.Errorf("task not found")
	}

	calculatedNode, exists := calculation.GetNodeByID(selectedExpr.AST, int(tr.NodeId))
	if !exists {
		return nil, fmt.Errorf("task not found")
	}

	if calculatedNode.Status != domain.StatusGivenToAgent {
		return nil, fmt.Errorf("task not found")
	}

	calculatedNode.Value = fmt.Sprintf("%.2f", tr.Result)
	calculatedNode.Left = nil
	calculatedNode.Right = nil
	calculatedNode.Status = domain.StatusFinished

	if calculatedNode == selectedExpr.AST {
		selectedExpr.Status = domain.StatusFinished
		selectedExpr.Result = tr.Result
		selectedExpr.AST = nil
	}

	return nil, nil
}

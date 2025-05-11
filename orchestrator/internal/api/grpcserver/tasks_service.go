package grpcserver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/bootstrap"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/repository"
	"github.com/SteeperMold/Calculator-go/orchestrator/pkg/calculation"
	pb "github.com/SteeperMold/Calculator-go/orchestrator/proto"
	"log"
	"strconv"
	"time"
)

type OrchestratorService struct {
	pb.UnimplementedOrchestratorServer
	repository           *repository.ExpressionsRepository
	contextTimeout       time.Duration
	timeAdditionMs       int
	timeSubtractionMs    int
	timeMultiplicationMs int
	timeDivisionMs       int
}

func NewOrchestratorService(db *sql.DB, config *bootstrap.Config) *OrchestratorService {
	return &OrchestratorService{
		repository:           repository.NewExpressionsRepository(db),
		contextTimeout:       config.ContextTimeout,
		timeAdditionMs:       config.TimeAdditionMs,
		timeSubtractionMs:    config.TimeSubtractionMs,
		timeMultiplicationMs: config.TimeMultiplicationMs,
		timeDivisionMs:       config.TimeDivisionMs,
	}
}

func (os *OrchestratorService) FetchTask(ctx context.Context, _ *pb.Empty) (*pb.GetTaskResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, os.contextTimeout)
	defer cancel()

	expr, err := os.repository.GetUnfinishedTask(ctx)
	if err != nil {
		if errors.Is(err, domain.ErrExpressionDoesntExist) {
			return nil, fmt.Errorf("no tasks available")
		}

		log.Printf("failed to get task: %v\n", err)
		return nil, fmt.Errorf("failed to get task")
	}

	deepestNode, exists := calculation.GetDeepestUnfinishedOperation(expr.AST)
	if !exists {
		return nil, fmt.Errorf("no tasks available")
	}

	arg1, err := strconv.ParseFloat(deepestNode.Left.Value, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to get task")
	}
	arg2, err := strconv.ParseFloat(deepestNode.Right.Value, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to get task")
	}

	operation := deepestNode.Value
	var operationTime int

	switch operation {
	case "+":
		operationTime = os.timeAdditionMs
	case "-":
		operationTime = os.timeSubtractionMs
	case "*":
		operationTime = os.timeMultiplicationMs
	case "/":
		operationTime = os.timeDivisionMs
	default:
		operationTime = 0
	}

	err = os.repository.SetNodeStatusGivenToAgent(ctx, deepestNode.Id)
	if err != nil {
		log.Printf("failed to get task: %v\n", err)
		return nil, fmt.Errorf("failed to get task")
	}

	deepestNode.Status = domain.StatusGivenToAgent

	return &pb.GetTaskResponse{
		ExpressionId:  expr.Id,
		NodeId:        deepestNode.Id,
		Arg1:          arg1,
		Arg2:          arg2,
		Operation:     deepestNode.Value,
		OperationTime: int32(operationTime),
	}, nil
}

func (os *OrchestratorService) SendResult(ctx context.Context, tr *pb.PostTaskResult) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, os.contextTimeout)
	defer cancel()

	err := os.repository.UpdateTaskById(ctx, tr.NodeId, tr.Result)
	if err != nil {
		log.Printf("failed to save result: %v\n", err)
		return nil, fmt.Errorf("failed to save result")
	}

	rootNodeId, err := os.repository.GetExpressionRootNodeId(ctx, tr.ExpressionId)
	if err != nil {
		log.Printf("failed to save result: %v\n", err)
		return nil, fmt.Errorf("failed to save result")
	}

	if tr.NodeId == *rootNodeId {
		err := os.repository.UpdateExpressionById(ctx, tr.ExpressionId, tr.Result)
		if err != nil {
			return nil, fmt.Errorf("failed to save result")
		}
	}

	return nil, nil
}

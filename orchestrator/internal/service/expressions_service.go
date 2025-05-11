package service

import (
	"context"
	"database/sql"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/repository"
	"github.com/SteeperMold/Calculator-go/orchestrator/pkg/calculation"
	"time"
)

type ExpressionsService struct {
	repository     *repository.ExpressionsRepository
	contextTimeout time.Duration
}

func NewExpressionsService(db *sql.DB, contextTimeout time.Duration) *ExpressionsService {
	return &ExpressionsService{
		repository:     repository.NewExpressionsRepository(db),
		contextTimeout: contextTimeout,
	}
}

func (es *ExpressionsService) CreateExpression(expression string) (*int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), es.contextTimeout)
	defer cancel()

	ast, err := calculation.BuildAST(expression)
	if err != nil {
		return nil, err
	}

	newExpr := &domain.Expression{
		Status: domain.StatusInProgress,
		Result: 0,
		AST:    ast,
	}

	return es.repository.CreateExpression(ctx, newExpr)
}

func (es *ExpressionsService) GetExpressionsList() ([]*domain.Expression, error) {
	ctx, cancel := context.WithTimeout(context.Background(), es.contextTimeout)
	defer cancel()

	return es.repository.GetExpressionsList(ctx)
}

func (es *ExpressionsService) GetExpressionById(id int64) (*domain.Expression, error) {
	ctx, cancel := context.WithTimeout(context.Background(), es.contextTimeout)
	defer cancel()

	return es.repository.GetExpressionById(ctx, id)
}

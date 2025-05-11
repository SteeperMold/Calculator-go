package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
)

func (er *ExpressionsRepository) GetUnfinishedTask(ctx context.Context) (*domain.Expression, error) {
	const q = `
		SELECT 
		    e.id, 
		    (SELECT s.name FROM statuses s WHERE s.id = e.status_id),
		    e.result 
		FROM expressions e
		WHERE e.status_id = (SELECT s.id FROM statuses s WHERE s.name = ?)
	 `

	row := er.db.QueryRowContext(ctx, q, domain.StatusInProgress)

	var expr domain.Expression
	err := row.Scan(&expr.Id, &expr.Status, &expr.Result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrExpressionDoesntExist
		}

		return nil, err
	}

	exprAST, err := er.loadNodes(ctx, expr.Id)
	if err != nil {
		return nil, err
	}
	expr.AST = exprAST

	return &expr, nil
}

func (er *ExpressionsRepository) SetNodeStatusGivenToAgent(ctx context.Context, nodeId int64) error {
	const q = `
		UPDATE nodes
		SET status_id = (SELECT s.id FROM statuses s WHERE s.name = ?)
		WHERE id = ?
	`

	_, err := er.db.ExecContext(ctx, q, domain.StatusGivenToAgent, nodeId)
	if err != nil {
		return err
	}

	return nil
}

func (er *ExpressionsRepository) UpdateTaskById(ctx context.Context, nodeId int64, result float64) error {
	const (
		deleteQ = `
			WITH node_ids AS (
				SELECT left_id, right_id
				FROM nodes
				WHERE id = ?
			)
	
			DELETE FROM nodes
			WHERE id IN (
				SELECT left_id FROM node_ids
				UNION 
				SELECT right_id FROM node_ids
			);
		`
		updateQ = `
			UPDATE nodes
			SET 
				status_id = (SELECT s.id FROM statuses s WHERE s.name = ?),
				value = ?,
				left_id = NULL,
				right_id = NULL
			WHERE id = ?;
		`
	)

	tx, err := er.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, deleteQ, nodeId)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, updateQ, domain.StatusFinished, result, nodeId)
	if err != nil {
		return nil
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (er *ExpressionsRepository) GetExpressionRootNodeId(ctx context.Context, exprId int64) (*int64, error) {
	const q = `
		SELECT id
		FROM nodes
		WHERE 
			expression_id = ? AND 
			parent_id IS NULL
	`

	row := er.db.QueryRowContext(ctx, q, exprId)

	var nodeId int64
	err := row.Scan(&nodeId)
	if err != nil {
		return nil, err
	}

	return &nodeId, nil
}

func (er *ExpressionsRepository) UpdateExpressionById(ctx context.Context, exprId int64, result float64) error {
	const (
		deleteQ = `
			DELETE FROM nodes
			WHERE 
			    expression_id = ? AND 
			    parent_id IS NULL
		`
		updateQ = `
			UPDATE expressions
			SET
				status_id = (SELECT s.id FROM statuses s WHERE s.name = ?),
				result = ?
			WHERE id = ?;
		`
	)

	tx, err := er.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, deleteQ, exprId)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, updateQ, domain.StatusFinished, result, exprId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

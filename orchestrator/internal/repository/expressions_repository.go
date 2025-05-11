package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	_ "github.com/mattn/go-sqlite3"
)

type ExpressionsRepository struct {
	db *sql.DB
}

func NewExpressionsRepository(db *sql.DB) *ExpressionsRepository {
	return &ExpressionsRepository{
		db: db,
	}
}

func (er *ExpressionsRepository) insertNode(ctx context.Context, exprId int64, parentId *int64, n *domain.Node) (int64, error) {
	const q = `
		INSERT INTO nodes(expression_id, parent_id, status_id, value) 
		VALUES(?, ?, (SELECT id FROM statuses WHERE name = ?), ?)
	`

	res, err := er.db.ExecContext(ctx, q, exprId, parentId, n.Status, n.Value)
	if err != nil {
		return 0, err
	}

	nodeId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if n.Left != nil {
		leftId, err := er.insertNode(ctx, exprId, &nodeId, n.Left)
		if err != nil {
			return 0, err
		}

		const q = `UPDATE nodes SET left_id = ? WHERE id = ?`

		_, err = er.db.ExecContext(ctx, q, leftId, nodeId)
		if err != nil {
			return 0, err
		}
	}

	if n.Right != nil {
		rightId, err := er.insertNode(ctx, exprId, &nodeId, n.Right)
		if err != nil {
			return 0, err
		}

		const q = `UPDATE nodes SET right_id = ? WHERE id = ?`

		_, err = er.db.ExecContext(ctx, q, rightId, nodeId)
		if err != nil {
			return 0, err
		}
	}

	return nodeId, nil
}

func (er *ExpressionsRepository) loadNodes(ctx context.Context, exprId int64) (*domain.Node, error) {
	const q = `
		SELECT 
		    n.id, 
		    n.parent_id, 
		    (SELECT s.name FROM statuses s WHERE s.id = n.status_id), 
		    n.value, 
		    n.left_id, 
		    n.right_id
		FROM nodes n
		WHERE n.expression_id = ?
	`

	rows, err := er.db.QueryContext(ctx, q, exprId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type rawNode struct {
		node    *domain.Node
		leftId  sql.NullInt64
		rightId sql.NullInt64
	}

	nodeMap := make(map[int64]*rawNode)
	var rootId int64

	for rows.Next() {
		var id int64
		var parentId, leftId, rightId sql.NullInt64
		var status, value string

		err := rows.Scan(&id, &parentId, &status, &value, &leftId, &rightId)
		if err != nil {
			return nil, err
		}

		nodeMap[id] = &rawNode{
			node: &domain.Node{
				Id:     id,
				Status: status,
				Value:  value,
			},
			leftId:  leftId,
			rightId: rightId,
		}

		if !parentId.Valid {
			rootId = id
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for _, rn := range nodeMap {
		if rn.leftId.Valid {
			rn.node.Left = nodeMap[rn.leftId.Int64].node
		}
		if rn.rightId.Valid {
			rn.node.Right = nodeMap[rn.rightId.Int64].node
		}
	}

	rootRaw, ok := nodeMap[rootId]
	if !ok {
		return nil, fmt.Errorf("root node not found")
	}

	return rootRaw.node, nil
}

func (er *ExpressionsRepository) CreateExpression(ctx context.Context, newExpr *domain.Expression) (*int64, error) {
	const q = `
		INSERT INTO expressions(status_id, result) 
		VALUES(
		   (SELECT id FROM statuses WHERE name = ?), 
		   ?
	   )
	`

	res, err := er.db.ExecContext(ctx, q, newExpr.Status, newExpr.Result)
	if err != nil {
		return nil, err
	}

	exprId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	_, err = er.insertNode(ctx, exprId, nil, newExpr.AST)

	return &exprId, nil
}

func (er *ExpressionsRepository) GetExpressionsList(ctx context.Context) ([]*domain.Expression, error) {
	const q = `
		SELECT 
		    e.id, 
		    (SELECT s.name FROM statuses s WHERE s.id = e.status_id),
		    e.result 
		FROM expressions e
	`

	rows, err := er.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exprs []*domain.Expression

	for rows.Next() {
		var expr domain.Expression

		err := rows.Scan(&expr.Id, &expr.Status, &expr.Result)
		if err != nil {
			return nil, err
		}

		if expr.Status != domain.StatusFinished {
			exprAST, err := er.loadNodes(ctx, expr.Id)
			if err != nil {
				return nil, err
			}
			expr.AST = exprAST
		}

		exprs = append(exprs, &expr)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exprs, nil
}

func (er *ExpressionsRepository) GetExpressionById(ctx context.Context, exprId int64) (*domain.Expression, error) {
	const q = `
		SELECT 
		    e.id, 
		    (SELECT s.name FROM statuses s WHERE s.id = e.status_id), 
		    e.result 
		FROM expressions e 
		WHERE e.id = ?
	`

	row := er.db.QueryRowContext(ctx, q, exprId)

	var expr domain.Expression
	err := row.Scan(&expr.Id, &expr.Status, &expr.Result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrExpressionDoesntExist
		}

		return nil, err
	}

	if expr.Status != domain.StatusFinished {
		exprAST, err := er.loadNodes(ctx, exprId)
		if err != nil {
			return nil, err
		}
		expr.AST = exprAST
	}

	return &expr, nil
}

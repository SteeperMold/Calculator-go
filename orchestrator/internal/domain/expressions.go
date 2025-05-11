package domain

import (
	"fmt"
)

const (
	StatusInProgress   = "in-progress"
	StatusGivenToAgent = "given-to-agent"
	StatusFinished     = "finished"
)

var (
	ErrInvalidExpression     = fmt.Errorf("invalid expression")
	ErrExpressionDoesntExist = fmt.Errorf("expression doesn't exist")
)

type Expression struct {
	Id     int64   `json:"id"`
	Status string  `json:"status"`
	Result float64 `json:"result"`
	AST    *Node
}

type Node struct {
	Id     int64
	Status string
	Value  string
	Left   *Node
	Right  *Node
}

type NewExpressionRequest struct {
	Expression string `json:"expression"`
}

type NewExpressionResponse struct {
	ID int64 `json:"id"`
}

type ExpressionsListResponse struct {
	Expressions []*Expression `json:"expressions"`
}

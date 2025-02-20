package domain

import "errors"

const (
	StatusInProgress   = "in-progress"
	StatusGivenToAgent = "given-to-agent"
	StatusFinished     = "finished"
)

var (
	ErrInvalidExpression = errors.New("invalid expression")
)

type Expression struct {
	ID     int     `json:"id"`
	Status string  `json:"status"`
	Result float64 `json:"result"`
	AST    *Node
}

type Node struct {
	ID     int
	Status string
	Value  string
	Left   *Node
	Right  *Node
}

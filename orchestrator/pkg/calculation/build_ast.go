package calculation

import (
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"strconv"
	"strings"
	"unicode"
)

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func tokenize(expr string) ([]string, error) {
	var tokens []string
	var num strings.Builder

	flushNumber := func() {
		if num.Len() > 0 {
			tokens = append(tokens, num.String())
			num.Reset()
		}
	}

	for _, r := range expr {
		switch {
		case unicode.IsSpace(r):
			flushNumber()
		case unicode.IsDigit(r) || r == '.':
			num.WriteRune(r)
		case r == '+' || r == '-' || r == '*' || r == '/' || r == '(' || r == ')':
			flushNumber()
			tokens = append(tokens, string(r))
		default:
			return nil, domain.ErrInvalidExpression
		}
	}
	flushNumber()
	return tokens, nil
}

func infixToRPN(expr string) ([]string, error) {
	tokens, err := tokenize(expr)
	if err != nil {
		return nil, err
	}

	var output []string
	var opStack []string

	for _, token := range tokens {
		if _, err := strconv.ParseFloat(token, 64); err == nil {
			output = append(output, token)
		} else if isOperator(token) {
			for len(opStack) > 0 && isOperator(opStack[len(opStack)-1]) &&
				precedence(opStack[len(opStack)-1]) >= precedence(token) {
				top := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				output = append(output, top)
			}
			opStack = append(opStack, token)
		} else if token == "(" {
			opStack = append(opStack, token)
		} else if token == ")" {
			foundParen := false
			for len(opStack) > 0 {
				top := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				if top == "(" {
					foundParen = true
					break
				}
				output = append(output, top)
			}
			if !foundParen {
				return nil, domain.ErrInvalidExpression
			}
		} else {
			return nil, domain.ErrInvalidExpression
		}
	}
	for len(opStack) > 0 {
		top := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		if top == "(" || top == ")" {
			return nil, domain.ErrInvalidExpression
		}
		output = append(output, top)
	}
	return output, nil
}

func BuildAST(expression string) (*domain.Node, error) {
	rpn, err := infixToRPN(expression)
	if err != nil {
		return nil, err
	}

	var stack []*domain.Node
	for _, token := range rpn {
		if _, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, &domain.Node{
				Value:  token,
				Status: domain.StatusInProgress,
			})
		} else if isOperator(token) {
			if len(stack) < 2 {
				return nil, domain.ErrInvalidExpression
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, &domain.Node{
				Value:  token,
				Left:   left,
				Right:  right,
				Status: domain.StatusInProgress,
			})
		} else {
			return nil, domain.ErrInvalidExpression
		}
	}
	if len(stack) != 1 {
		return nil, domain.ErrInvalidExpression
	}
	return stack[0], nil
}

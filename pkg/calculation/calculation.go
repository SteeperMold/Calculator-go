package calculation

import (
	"strings"
	"unicode"
)

var precedence = map[rune]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

func isExpressionValid(expression string) bool {
	if strings.TrimSpace(expression) == "" {
		return false
	}

	for _, r := range expression {
		if !strings.ContainsRune("0123456789+-*/() ", r) {
			return false
		}
	}

	return isParenthesisSequenceCorrect(expression)
}

func isParenthesisSequenceCorrect(expression string) bool {
	var stack []rune

	for _, r := range expression {
		if r == '(' {
			stack = append(stack, r)
		} else if r == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func applyOperator(operators []rune, values []float64) ([]rune, []float64, error) {
	if len(values) < 2 || len(operators) == 0 {
		return operators, values, ErrInvalidExpression
	}

	operator := operators[len(operators)-1]
	left := values[len(values)-2]
	right := values[len(values)-1]

	operators = operators[:len(operators)-1]
	values = values[:len(values)-2]

	var result float64
	switch operator {
	case '+':
		result = left + right
	case '-':
		result = left - right
	case '*':
		result = left * right
	case '/':
		if right == 0 {
			return operators, values, ErrInvalidExpression
		}
		result = left / right
	default:
		return operators, values, ErrInvalidExpression
	}

	values = append(values, result)
	return operators, values, nil
}

func Calculate(expression string) (float64, error) {
	if !isExpressionValid(expression) {
		return 0, ErrInvalidExpression
	}

	var operators []rune
	var values []float64

	for index := 0; index < len(expression); index++ {
		r := rune(expression[index])

		if r == ' ' {
			continue
		}

		if unicode.IsDigit(r) {
			number := 0
			for index < len(expression) && unicode.IsDigit(rune(expression[index])) {
				number = number*10 + int(expression[index]-'0')
				index++
			}
			index--
			values = append(values, float64(number))
		} else if r == '(' {
			operators = append(operators, r)
		} else if r == ')' {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				var err error
				operators, values, err = applyOperator(operators, values)
				if err != nil {
					return 0, err
				}
			}

			if len(operators) == 0 {
				return 0, ErrInvalidExpression
			}
			operators = operators[:len(operators)-1]
		} else if strings.ContainsRune("+-/*", r) {
			for len(operators) > 0 &&
				operators[len(operators)-1] != '(' &&
				precedence[operators[len(operators)-1]] >= precedence[r] {
				var err error
				operators, values, err = applyOperator(operators, values)
				if err != nil {
					return 0, err
				}
			}
			operators = append(operators, r)
		} else {
			return 0, ErrInvalidExpression
		}
	}

	for len(operators) > 0 {
		var err error
		operators, values, err = applyOperator(operators, values)
		if err != nil {
			return 0, err
		}
	}

	if len(values) != 1 {
		return 0, ErrInvalidExpression
	}

	return values[0], nil
}

package main

import (
	"fmt"
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
	for _, r := range expression {
		if !strings.ContainsRune("0123456789+-*/() ", r) {
			return false
		}
	}

	return true
}

func isParenthesisSequenceCorrect(expression string) bool {
	var stack []int

	for index, r := range expression {
		if r == '(' {
			stack = append(stack, index)
		} else if r == ')' {
			if len(stack) == 0 {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func applyOperator(operators []rune, values []float64) ([]rune, []float64) {
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
		result = left / right
	}

	values = append(values, result)
	return operators, values
}

func Calc(expression string) (float64, error) {
	if !isExpressionValid(expression) {
		return 0, fmt.Errorf("bad expression string")
	}

	if !isParenthesisSequenceCorrect(expression) {
		return 0, fmt.Errorf("wrong parenthesis sequence")
	}

	if strings.TrimSpace(expression) == "" {
		return 0, fmt.Errorf("empty expression")
	}

	var operators []rune
	var values []float64

	for index, r := range expression {
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
			for operators[len(operators)-1] != '(' {
				operators, values = applyOperator(operators, values)
			}

			operators = operators[:len(operators)-1]
		} else if strings.ContainsRune("+-/*", r) {
			for len(operators) > 0 &&
				operators[len(operators)-1] != '(' &&
				precedence[operators[len(operators)-1]] >= precedence[r] {
				operators, values = applyOperator(operators, values)
			}

			operators = append(operators, r)
		}
	}

	// Если операторов столько сколько нужно, то останется два числа и один оператор
	if !(len(values) > len(operators)) {
		return 0, fmt.Errorf("invalid operators")
	}

	for len(operators) > 0 {
		operators, values = applyOperator(operators, values)
	}

	return values[0], nil
}

func main() {
	res, err := Calc("2 + 2")

	if err != nil {
		fmt.Printf("failew with error: %v\n", err)
	}

	fmt.Printf("res = %v\n", res)
}

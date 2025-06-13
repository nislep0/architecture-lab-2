package lab2

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func PostfixToInfix(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", errors.New("input is empty")
	}

	tokens := strings.Fields(input)
	stack := []string{}

	isOperator := func(token string) bool {
		return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
	}

	for _, token := range tokens {
		// якщо операнд (може бути багатозначне число)
		if isOperator(token) {
			// ѕотр≥бно два операнди в стеку
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression: not enough operands for operator '%s'", token)
			}
			// ¬ит€гуЇмо два операнди з к≥нц€ стека
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// ‘ормуЇмо новий вираз
			expr := fmt.Sprintf("(%s %s %s)", left, token, right)
			stack = append(stack, expr)
		} else if isValidOperand(token) {
			stack = append(stack, token)
		} else {
			return "", fmt.Errorf("invalid token: %s", token)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("invalid expression: extra operands or missing operators")
	}

	return stack[0], nil
}

func isValidOperand(token string) bool {
	// ƒозволимо числов≥ значенн€ з дес€тковою крапкою
	for i, r := range token {
		if !(unicode.IsDigit(r) || r == '.') || (r == '.' && (i == 0 || i == len(token)-1)) {
			return false
		}
	}
	return true
}

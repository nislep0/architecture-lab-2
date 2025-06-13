package lab2

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)
// PrefixToPostfix перетворює префіксний вираз у постфіксний.
// Підтримуються оператори +, -, *, /, ^.
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
		
		if isOperator(token) {
			
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression: not enough operands for operator '%s'", token)
			}
			
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			
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
	
	for i, r := range token {
		if !(unicode.IsDigit(r) || r == '.') || (r == '.' && (i == 0 || i == len(token)-1)) {
			return false
		}
	}
	return true
}

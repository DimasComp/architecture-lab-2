package lab2

import (
	"fmt"
	"strings"
)

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "^"
}

func PostfixToInfix(postfix string) (string, error) {
	stack := []string{}
	tokens := strings.Fields(postfix)

	for _, token := range tokens {
		if !isOperator(token) {
			stack = append(stack, token)
		} else {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid postfix expression")
			}
			lastIndex := len(stack) - 1

			op2 := stack[lastIndex]
			stack = stack[:lastIndex]
			lastIndex--
			op1 := stack[lastIndex]
			stack = stack[:lastIndex]

			infix := fmt.Sprintf("%s %s %s", op1, token, op2)
			stack = append(stack, infix)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid postfix expression")
	}

	return stack[0], nil
}

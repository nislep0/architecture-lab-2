package lab2

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleExpressions(t *testing.T) {
	result, err := PostfixToInfix("3 4 +")
	assert.NoError(t, err)
	assert.Equal(t, "(3 + 4)", result)

	result, err = PostfixToInfix("10 2 /")
	assert.NoError(t, err)
	assert.Equal(t, "(10 / 2)", result)

	result, err = PostfixToInfix("5 1 2 + 4 * + 3 -")
	assert.NoError(t, err)
	assert.Equal(t, "((5 + ((1 + 2) * 4)) - 3)", result)
}

func TestComplexExpression(t *testing.T) {
	expr := "2 3 + 5 * 6 2 / - 4 2 ^ +"
	result, err := PostfixToInfix(expr)
	assert.NoError(t, err)
	assert.Equal(t, "((((2 + 3) * 5) - (6 / 2)) + (4 ^ 2))", result)
}

func TestEmptyInput(t *testing.T) {
	_, err := PostfixToInfix("")
	assert.Error(t, err)
}

func TestExtraOperands(t *testing.T) {
	_, err := PostfixToInfix("3 4 5 +")
	assert.Error(t, err)
}

func TestTooFewOperands(t *testing.T) {
	_, err := PostfixToInfix("+")
	assert.Error(t, err)
}

func TestInvalidToken(t *testing.T) {
	_, err := PostfixToInfix("3 4 &")
	assert.Error(t, err)
}

func ExamplePostfixToInfix() {
	result, _ := PostfixToInfix("2 3 + 4 *")
	fmt.Println(result)
	// Output: ((2 + 3) * 4)
}
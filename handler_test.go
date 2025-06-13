package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	)

func TestComputeHandler_ValidPostfix(t *testing.T) {
	input := strings.NewReader("3 4 +\n")
	output := &bytes.Buffer{}

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	require.NoError(t, err)
	assert.Equal(t, "(3 + 4)\n", output.String())
}

func TestComputeHandler_InvalidPostfix(t *testing.T) {
	input := strings.NewReader("3 +\n")
	output := &bytes.Buffer{}

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid")
}

func TestComputeHandler_EmptyInput(t *testing.T) {
	input := strings.NewReader("\n")
	output := &bytes.Buffer{}

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "empty")
}

func TestComputeHandler_ReadsOnlyOneLine(t *testing.T) {
	input := strings.NewReader("2 2 *\n5 1 -\n")
	output := &bytes.Buffer{}

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	require.NoError(t, err)
	assert.Equal(t, "(2 * 2)\n", output.String())
}

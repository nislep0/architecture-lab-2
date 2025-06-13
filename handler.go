package lab2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.Input)

	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read input: %w", err)
	}

	line = strings.TrimSpace(line)
	if line == "" {
		return fmt.Errorf("empty input expression")
	}

	result, err := PostfixToInfix(line)
	if err != nil {
		return fmt.Errorf("failed to convert expression: %w", err)
	}

	_, err = fmt.Fprintln(ch.Output, result)
	if err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	return nil
}
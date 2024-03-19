package lab2

import (
	"bufio"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.Input)

	data, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	result, err := PostfixToInfix(data)
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(result))
	if err != nil {
		return err
	}

	return nil
}

package lab2_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DimasComp/architecture-lab-2"
)

func TestComputeHandler(t *testing.T) {

	t.Run("Write test", func(t *testing.T) {
        input := "2 3 + 4 $$ 5 * -" //broke that line so test fails
        expected := "2 + 3 - 4 * 5"
		buf := bytes.NewBuffer([]byte{})

		handler := lab2.ComputeHandler{Input: strings.NewReader(input), Output: buf}
		err := handler.Compute()
		result := buf.String()
        assert.NoError(t, err)
        assert.Equal(t, expected, result)
    })

	t.Run("Error test", func(t *testing.T) {
        input := "2 3 2 $"
		buf := bytes.NewBuffer([]byte{})

		handler := lab2.ComputeHandler{Input: strings.NewReader(input), Output: buf}
		err := handler.Compute()
        assert.Error(t, err)
    })
}

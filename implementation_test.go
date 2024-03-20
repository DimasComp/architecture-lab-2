package lab2_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/DimasComp/architecture-lab-2"
)

func TestPostfixToInfix(t *testing.T) {
    t.Run("Three operations", func(t *testing.T) {
        input := "2 3 + 4 5 * -"
        expected := "2 + 3 - 4 * 5"
        result, err := lab2.PostfixToInfix(input)
        assert.NoError(t, err)
        assert.Equal(t, expected, result)
    })

	t.Run("Four operations", func(t *testing.T) {
        input := "1 3 2 * + 6 2 / -"
        expected := "1 + 3 * 2 - 6 / 2"
        result, err := lab2.PostfixToInfix(input)
        assert.NoError(t, err)
        assert.Equal(t, expected, result)
    })

	t.Run("Eight operations", func(t *testing.T) {
        input := "1 2 + 3 4 * 5 / - 6 - 7 8 / - 9 +"
        expected := "1 + 2 - 3 * 4 / 5 - 6 - 7 / 8 + 9"
        result, err := lab2.PostfixToInfix(input)
        assert.NoError(t, err)
        assert.Equal(t, expected, result)
    })

    t.Run("Just one number", func(t *testing.T) {
        input := "7"
        expected := "7"
        result, err := lab2.PostfixToInfix(input)
        assert.NoError(t, err)
        assert.Equal(t, expected, result)
    })

    t.Run("Invalid characters", func(t *testing.T) {
        input := "2 @ 3 +"
        _, err := lab2.PostfixToInfix(input)
        assert.Error(t, err)
    })
}
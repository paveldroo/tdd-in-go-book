package calculator_test

import (
	"github.com/paveldroo/tdd-in-go-book/calculator/calculator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	actAssert := func(x, y, want float64) {
		// Act
		got := e.Add(x, y)

		// Assert
		if got != want {
			t.Errorf("Add(%.2f, %.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
		}
	}

	t.Run("positive input", func(t *testing.T) {
		x, y := 2.5, 3.5
		want := 6.0
		actAssert(x, y, want)
	})

	t.Run("negative input", func(t *testing.T) {
		x, y := -2.5, -3.5
		want := -6.0
		actAssert(x, y, want)
	})
}

func TestEngine_ProcessOperation(t *testing.T) {
	t.Run("valid operation", func(t *testing.T) {
		// Arrange
		operation := calculator.Operation{
			Expression: "2 + 3",
			Operator:   "+",
			Operands:   []float64{2.0, 3.0},
		}
		engine := calculator.NewEngine()
		expectedRes := "2 + 3 = 5.00"

		// Act
		res, err := engine.ProcessOperation(operation)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, res)
		assert.Contains(t, *res, expectedRes)
	})
	t.Run("invalid operation", func(t *testing.T) {
		// Arrange
		operation := calculator.Operation{
			Expression: "2 % 3",
			Operator:   "%",
			Operands:   []float64{2.0, 3.0},
		}
		engine := calculator.NewEngine()
		expectedErr := "unknown operator in expr"

		// Act
		res, err := engine.ProcessOperation(operation)

		// Assert
		require.Nil(t, res)
		require.NotNil(t, err)
		assert.Contains(t, err.Error(), expectedErr)
	})
}

func BenchmarkAdd(b *testing.B) {
	e := calculator.Engine{}

	// run the Add function b.N times
	for i := 0; i < b.N; i++ {
		e.Add(2, 3)
	}
}

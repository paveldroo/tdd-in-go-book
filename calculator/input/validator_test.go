package input_test

import (
	"github.com/paveldroo/tdd-in-go-book/calculator/input"
	"testing"
)

func TestValidator_CheckInput(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		// Arrange
		validator := setup(t, []string{"+"})
		operator := "+"
		operands := []float64{2.0, 3.0}

		// Act
		err := validator.CheckInput(operator, operands)

		// Assert
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("invalid operands quantity", func(t *testing.T) {
		// Arrange
		validator := setup(t, []string{"+"})
		operator := "+"
		operands := []float64{3.0}

		// Act
		err := validator.CheckInput(operator, operands)

		// Assert
		if err == nil {
			t.Fatal(err)
		}
	})
	t.Run("invalid operator", func(t *testing.T) {
		// Arrange
		validator := setup(t, []string{"+"})
		operator := "-"
		operands := []float64{2.0, 3.0}

		// Act
		err := validator.CheckInput(operator, operands)

		// Assert
		if err == nil {
			t.Fatal(err)
		}
	})

}

func setup(t *testing.T, validOps []string) *input.Validator {
	t.Helper()
	return input.NewValidator(2, validOps)
}

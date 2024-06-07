package input_test

import (
	"fmt"
	"github.com/paveldroo/tdd-in-go-book/calculator"
	"github.com/paveldroo/tdd-in-go-book/input"
	"github.com/paveldroo/tdd-in-go-book/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser_ProcessExpression(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		// Arrange
		expr := "2 + 3"
		operator := "+"
		operands := []float64{2.0, 3.0}
		expectedResult := "2 + 3 = 5.5"
		engine := mocks.NewOperationProcessor(t)
		validator := mocks.NewValidationHelper(t)
		parser := input.NewParser(engine, validator)

		validator.On("CheckInput", operator, operands).Return(nil).Once()
		engine.On("ProcessOperation", &calculator.Operation{
			Expression: expr,
			Operator:   operator,
			Operands:   operands,
		}).Return(expectedResult, nil).Once()

		// Act
		result, err := parser.ProcessExpression(expr)

		// Assert
		require.Nil(t, err)
		assert.Equal(t, expectedResult, result)

		// Other assertions
		validator.AssertExpectations(t)
		engine.AssertExpectations(t)
	})
	t.Run("invalid operation", func(t *testing.T) {
		// Arrange
		expr := "2 % 3"
		operator := "%"
		operands := []float64{2.0, 3.0}
		expectedErrMsg := "bad operator"
		engine := mocks.NewOperationProcessor(t)
		validator := mocks.NewValidationHelper(t)
		parser := input.NewParser(engine, validator)
		validator.On("CheckInput", operator, operands).Return(fmt.Errorf(expectedErrMsg)).Once()

		// Act
		result, err := parser.ProcessExpression(expr)

		// Assert
		require.NotNil(t, err)
		assert.Equal(t, "", result)
		assert.Contains(t, err.Error(), expr)
		assert.Contains(t, err.Error(), expectedErrMsg)
		validator.AssertExpectations(t)
	})
}

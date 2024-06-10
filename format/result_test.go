package format_test

import (
	"fmt"
	"github.com/paveldroo/tdd-in-go-book/format"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResult(t *testing.T) {
	// Arrange
	expr := "2 + 3"
	result := 5.55

	// Act
	resultStr := format.Result(expr, result)

	// Assert
	assert.Contains(t, resultStr, expr)
	assert.Contains(t, resultStr, fmt.Sprint(result))
}

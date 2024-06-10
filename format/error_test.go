package format_test

import (
	"fmt"
	"github.com/paveldroo/tdd-in-go-book/format"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestError(t *testing.T) {
	// Arrange
	expr := "abc + 3"
	expectedErr := "very bad error"

	// Act
	err := format.Error(expr, fmt.Errorf(expectedErr))

	// Assert
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), expr)
	assert.Contains(t, err.Error(), expectedErr)
}

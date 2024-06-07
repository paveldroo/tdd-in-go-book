package format

import "fmt"

func Result(expression string, result float64) string {
	// implementation code
	return ""
}

func Error(expr string, err error) error {
	return fmt.Errorf("invalid expression: %v, error: %v", expr, err)
}

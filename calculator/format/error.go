package format

import "fmt"

func Error(expr string, err error) error {
	return fmt.Errorf("CALCULATION FAILED: expression %s is invalid: %v", expr, err)
}

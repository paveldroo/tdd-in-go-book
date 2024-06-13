package input

import "fmt"

type Validator struct {
	expectedLength int
	validOperators []string
}

func NewValidator(len int, validOps []string) *Validator {
	return &Validator{
		expectedLength: len,
		validOperators: validOps,
	}
}

func (v *Validator) CheckInput(op string, operands []float64) error {
	if len(operands) != v.expectedLength {
		return fmt.Errorf("invalid operands qiantity: %v", operands)
	}

	err := v.CheckOperator(op)
	if err != nil {
		return err
	}

	return nil
}

func (v *Validator) CheckOperator(op string) error {
	for _, vops := range v.validOperators {
		if vops == op {
			return nil
		}
	}
	return fmt.Errorf("invalid operator: %v", op)
}

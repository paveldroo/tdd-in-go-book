package input

import "fmt"

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) CheckInput(op string, operands []float64) error {
	if len(op) != 1 {
		return fmt.Errorf("bad operator")
	}

	if op != "+" && op != "-" && op != "/" && op != "*" {
		return fmt.Errorf("bad operator")
	}

	if len(operands) != 2 {
		return fmt.Errorf("wrong quantity of operands")
	}

	return nil
}

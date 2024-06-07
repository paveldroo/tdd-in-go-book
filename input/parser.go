package input

import (
	"fmt"
	"github.com/paveldroo/tdd-in-go-book/calculator"
	"github.com/paveldroo/tdd-in-go-book/format"
	"strconv"
	"strings"
)

type OperationProcessor interface {
	ProcessOperation(operation *calculator.Operation) (string, error)
}

type ValidationHelper interface {
	CheckInput(operator string, operands []float64) error
}

type Parser struct {
	engine    OperationProcessor
	validator ValidationHelper
}

func NewParser(e OperationProcessor, v ValidationHelper) *Parser {
	return &Parser{
		engine:    e,
		validator: v,
	}
}

func (p *Parser) ProcessExpression(expr string) (string, error) {
	op, err := p.getOperation(expr)
	if err != nil {
		return "", format.Error(expr, err)
	}
	return p.engine.ProcessOperation(op)
}

func (p *Parser) getOperation(expr string) (*calculator.Operation, error) {
	op := calculator.Operation{Expression: expr}

	fields := strings.Fields(expr)

	if len(fields) != 3 {
		return nil, fmt.Errorf("got invalid expression: %v", expr)
	}

	operand, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return nil, fmt.Errorf("got invalid expression: %v", expr)
	}
	op.Operands = append(op.Operands, operand)

	operand, err = strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return nil, fmt.Errorf("got invalid expression: %v", expr)
	}
	op.Operands = append(op.Operands, operand)

	op.Operator = fields[1]

	if err = p.validator.CheckInput(op.Operator, op.Operands); err != nil {
		return nil, err
	}

	return &op, nil
}

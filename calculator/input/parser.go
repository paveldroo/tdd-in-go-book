package input

import (
	"fmt"
	"github.com/paveldroo/tdd-in-go-book/calculator/calculator"
	"github.com/paveldroo/tdd-in-go-book/calculator/format"
	"strconv"
	"strings"
)

type OperationProcessor interface {
	ProcessOperation(operation calculator.Operation) (*string, error)
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

func (p *Parser) ProcessExpression(expr string) (*string, error) {
	op, err := p.getOperation(expr)
	if err != nil {
		return nil, format.Error(expr, err)
	}
	return p.engine.ProcessOperation(*op)
}

func (p *Parser) getOperation(expr string) (*calculator.Operation, error) {
	ops := strings.Fields(expr)

	if len(ops) != 3 {
		return nil, fmt.Errorf("invalid expression length: %v", expr)
	}

	leftOp, err := strconv.ParseFloat(ops[0], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid left operand: %v", ops[0])
	}

	rightOp, err := strconv.ParseFloat(ops[2], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid right operand: %v", ops[2])
	}

	operator := ops[1]
	operands := []float64{leftOp, rightOp}
	if err = p.validator.CheckInput(operator, operands); err != nil {
		return nil, err
	}

	return &calculator.Operation{
		Expression: expr,
		Operator:   operator,
		Operands:   operands,
	}, nil
}

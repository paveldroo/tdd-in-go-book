package input

import (
	"fmt"
	"github.com/paveldroo/tdd-in-go-book/calculator"
	"github.com/paveldroo/tdd-in-go-book/format"
	"strconv"
	"strings"
)

type OperationProcessor interface {
	ProcessOperation(operation *calculator.Operation) (*string, error)
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
	return p.engine.ProcessOperation(op)
}

func (p *Parser) getOperation(expr string) (*calculator.Operation, error) {
	op := calculator.Operation{Expression: expr}

	operator := ""
	switch {
	case strings.Contains(expr, "+"):
		operator = "+"
	case strings.Contains(expr, "-"):
		operator = "-"
	case strings.Contains(expr, "*"):
		operator = "*"
	case strings.Contains(expr, "/"):
		operator = "/"
	default:
		return nil, fmt.Errorf("got invalid expression: %v", expr)
	}

	op.Operator = operator

	splits := strings.Split(expr, operator)
	if len(splits) != 2 {
		return nil, fmt.Errorf("got invalid expression: %v", expr)
	}
	fo, err := strconv.ParseFloat(strings.TrimSpace(splits[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("got invalid expression: %v", expr)
	}
	so, err := strconv.ParseFloat(strings.TrimSpace(splits[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("got invalid expression: %v", expr)
	}
	op.Operands = []float64{fo, so}

	return &op, nil
}

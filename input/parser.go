package input

import (
	"github.com/paveldroo/tdd-in-go-book/calculator"
	"github.com/paveldroo/tdd-in-go-book/format"
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
	return nil, nil
}

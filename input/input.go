package input

import "github.com/paveldroo/tdd-in-go-book/calculator"

type Parser struct {
	engine    *calculator.Engine
	validator *Validator
}

func (p *Parser) ProcessExpression(expression string) (*string, error) {
	// implementation code
}

// method declarations

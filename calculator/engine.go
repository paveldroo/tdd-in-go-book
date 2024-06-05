package calculator

import (
	"fmt"
	"strconv"
)

type Operation struct {
	Expression string
	Operator   string
	Operands   []float64
}

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) ProcessOperation(op *Operation) (*string, error) {
	switch {
	case op.Operator == "+":
		res := e.Add(op.Operands[0], op.Operands[1])
		resStr := strconv.FormatFloat(res, 'f', 6, 64)
		return &resStr, nil
	case op.Operator == "-":
		res := e.Add(op.Operands[0], op.Operands[1])
		resStr := strconv.FormatFloat(res, 'f', 6, 64)
		return &resStr, nil
	case op.Operator == "*":
		res := e.Add(op.Operands[0], op.Operands[1])
		resStr := strconv.FormatFloat(res, 'f', 6, 64)
		return &resStr, nil
	case op.Operator == "/":
		res := e.Add(op.Operands[0], op.Operands[1])
		resStr := strconv.FormatFloat(res, 'f', 6, 64)
		return &resStr, nil
	}

	return nil, fmt.Errorf("unknown operator in expr %v", op.Expression)
}

func (e *Engine) Add(x, y float64) float64 {
	return x + y
}

type Calculator struct {
	engine *Engine
}

func NewCalculator(e *Engine) *Calculator {
	return &Calculator{engine: e}
}

func (c Calculator) PrintAdd(x, y float64) {
	fmt.Println("Result:", c.engine.Add(x, y))
}

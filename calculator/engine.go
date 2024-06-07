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

func (e *Engine) ProcessOperation(op Operation) (*string, error) {
	opResStr := ""
	switch {
	case op.Operator == "+":
		opRes := e.Add(op.Operands[0], op.Operands[1])
		opResStr = strconv.FormatFloat(opRes, 'f', 6, 64)
	case op.Operator == "-":
		res := e.Subtract(op.Operands[0], op.Operands[1])
		opResStr = strconv.FormatFloat(res, 'f', 6, 64)
	case op.Operator == "*":
		res := e.Multiply(op.Operands[0], op.Operands[1])
		opResStr = strconv.FormatFloat(res, 'f', 6, 64)
	case op.Operator == "/":
		res := e.Divide(op.Operands[0], op.Operands[1])
		opResStr = strconv.FormatFloat(res, 'f', 6, 64)
	default:
		return nil, fmt.Errorf("unknown operator in expr %v", op.Expression)
	}

	resStr := fmt.Sprintf("%v %v %v = %v", op.Operands[0], op.Operator, op.Operands[1], opResStr)

	return &resStr, nil
}

func (e *Engine) Add(x, y float64) float64 {
	return x + y
}

func (e *Engine) Subtract(x, y float64) float64 {
	return x - y
}

func (e *Engine) Multiply(x, y float64) float64 {
	return x * y
}

func (e *Engine) Divide(x, y float64) float64 {
	return x / y
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

package calculator

import "fmt"

type Operation struct {
	Expression string
	Operator   string
	Operands   []float64
}

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e Engine) Add(x, y float64) float64 {
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

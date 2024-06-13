package main

import (
	"github.com/paveldroo/tdd-in-go-book/calculator/calculator"
)

func main() {
	e := calculator.NewEngine()
	calc := calculator.NewCalculator(e)
	calc.PrintAdd(2.5, 6.3)
}

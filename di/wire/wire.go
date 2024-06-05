//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"
import "github.com/paveldroo/tdd-in-go-book/calculator"

var Set = wire.NewSet(calculator.NewEngine, wire.Bind(new(calculator.Adder), new(*calculator.Engine)), calculator.NewCalculator)

func InitCalc() *calculator.Calculator {
	wire.Build(Set)
	return nil
}

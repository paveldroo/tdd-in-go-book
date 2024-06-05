// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/paveldroo/tdd-in-go-book/calculator"
)

// Injectors from wire.go:

func InitCalc() *calculator.Calculator {
	engine := calculator.NewEngine()
	calculatorCalculator := calculator.NewCalculator(engine)
	return calculatorCalculator
}

// wire.go:

var Set = wire.NewSet(calculator.NewEngine, wire.Bind(new(calculator.Adder), new(*calculator.Engine)), calculator.NewCalculator)

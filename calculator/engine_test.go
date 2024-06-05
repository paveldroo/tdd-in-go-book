package calculator_test

import (
	"github.com/paveldroo/tdd-in-go-book/calculator"
	"log"
	"os"
	"testing"
)

func init() {
	log.Println("Init setup.")
}

func TestMain(m *testing.M) {
	// setup statements
	setup()

	// run the tests
	e := m.Run()

	// cleanup statements
	teardown()

	// report the exit code
	os.Exit(e)
}

func setup() {
	log.Println("Setting up.")
}

func teardown() {
	log.Println("Tearing down.")
}

func TestAdd(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	// Arrange
	e := calculator.Engine{}
	x, y := 2.5, 3.5
	want := 6.0

	// Act
	got := e.Add(x, y)

	// Assert
	if got != want {
		t.Errorf("Add(%.2f, %.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}

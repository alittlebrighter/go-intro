package main

import (
	"fmt"

	"github.com/alittlebrighter/go-intro/calculator"
)

func main() {
	var calc calculator.Calculator
	var err error

	calc, err = calculator.NewCalculator()
	if err != nil {
		panic("could not create calculator")
	}

	op := calculator.Add
	parameters := []int{4, 1, 2}

	// CalcOperation implements String() so format uses that automatically
	fmt.Printf("Applying %s operation on %v equals %d\n", op, parameters, calc.SetParams(parameters...).Apply(op))
}

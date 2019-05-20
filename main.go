package main

import (
	"github.com/alittlebrighter/go-intro/calculator"
	"github.com/alittlebrighter/go-intro/io"
)

func main() {
	var calc calculator.Calculator
	var err error

	calc, err = calculator.NewCalculator()
	if err != nil {
		panic("could not create calculator")
	}

	in, out := make(chan io.Msg), make(chan io.Msg)

	go runCalculator(calc, in, out)

	io.StartServer("0.0.0.0:9000", in, out)
}

func runCalculator(calc calculator.Calculator, requests <-chan io.Msg, responses chan io.Msg) {
	for req := range requests {
		req.Result = calc.SetParams(req.Params...).Apply(req.Operation)
		responses <- req
	}
}

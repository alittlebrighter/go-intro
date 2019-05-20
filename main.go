package main

import (
	"log"
	"runtime"
	"time"

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

	cpus := runtime.NumCPU()

	in, out := make(chan io.Msg, cpus), make(chan io.Msg, cpus)

	for i := 0; i < cpus; i++ {
		go runCalculator(calc, in, out)
	}
	log.Printf("running %d workers", cpus)

	io.StartServer("0.0.0.0:9000", in, out)
}

func runCalculator(calc calculator.Calculator, requests <-chan io.Msg, responses chan io.Msg) {
	for req := range requests {
		req.Result = calc.SetParams(req.Params...).Apply(req.Operation)
		time.Sleep(time.Duration(req.Result) * time.Second)
		responses <- req
	}
}

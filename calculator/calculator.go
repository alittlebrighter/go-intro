package calculator

import (
	"errors"
	"strings"
)

type CalcOperation uint8

const ( // effectively an enum
	Add CalcOperation = iota
	Subtract
	Multiply
	Divide
)

func (op CalcOperation) IntApply(params ...int) int {
	var result int
	if len(params) == 0 {
		return result
	} else if len(params) == 1 {
		return params[0]
	}

	switch op {
	case Add:
		result = params[0] + params[1]
	case Subtract:
		result = params[0] - params[1]
	case Multiply:
		result = params[0] * params[1]
	case Divide:
		result = params[0] / params[1]
	}

	if len(params) > 2 {
		return op.IntApply(append([]int{result}, params[2:]...)...) // recursion
	} else {
		return result
	}
}

var opStrings = []string{"Add", "Subtract", "Multiply", "Divide"}

func (op CalcOperation) String() string {
	return opStrings[op]
}

func ParseCalcOperation(opStr string) (op CalcOperation, err error) { // named and initialized return variables
	switch strings.ToLower(opStr) {
	case "add":
		fallthrough // note default is to `break` at the end of a `case` so we need to explicitly `fallthrough`
	case "+":
		op = Add
	case "subtract":
		fallthrough
	case "-":
		op = Subtract
	case "multiply":
		fallthrough
	case "*":
		op = Multiply
	case "divide":
		fallthrough
	case "/":
		op = Divide
	default:
		err = errors.New("unknown operation")
	}
	return
}

type Calculator interface {
	SetParams(...int) Calculator
	Apply(CalcOperation) int
}

type calculator struct {
	params []int
}

func (calc *calculator) SetParams(params ...int) Calculator {
	calc.params = params
	return calc
}

func (calc *calculator) Apply(op CalcOperation) int {
	return op.IntApply(calc.params...)
}

func NewCalculator() (*calculator, error) {
	return (&calculator{params: []int{}}), nil
}

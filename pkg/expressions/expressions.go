package expressions

import (
	"errors"
	"fmt"
)

type ExprRunner struct {
	A          bool     `json:"a"`
	B          bool     `json:"b"`
	C          bool     `json:"c"`
	D          float64  `json:"d"`
	E          int      `json:"e"`
	F          int      `json:"f"`
	Options    []string `json:"options"`
	optionsMap map[string]int
	H          float64
	K          float64
}

const M, P, T = 1.0, 2.0, 3.0
const (
	OptionBase    = "BASE"
	OptionCustom1 = "CUSTOM1"
	OptionCustom2 = "CUSTOM2"
)

var expressionOptions = map[string]int{
	OptionBase:    0,
	OptionCustom1: 1,
	OptionCustom2: 2,
}

func (e *ExprRunner) Run() error {
	e.optionsToMap()

	err := e.validOptions()
	if err != nil {
		return err
	}

	if e.hasOption(OptionBase) {
		err = e.evalBaseExpression()
		if err != nil {
			return err
		}
	}

	if e.hasOption(OptionCustom1) {
		e.evalCustomExpression1()
	}

	if e.hasOption(OptionCustom2) {
		e.evalCustomExpression2()
	}

	return nil
}

func (e *ExprRunner) optionsToMap() {
	e.optionsMap = make(map[string]int)
	for _, option := range e.Options {
		e.optionsMap[option]++
	}
}

func (e ExprRunner) validOptions() error {
	if len(e.optionsMap) == 0 {
		return errors.New("no expression options found")
	}

	for option, count := range e.optionsMap {
		if count > 1 {
			return fmt.Errorf("option %q used multiple times", option)
		}

		if _, ok := expressionOptions[option]; !ok {
			return fmt.Errorf("%q is not a valid expression option", option)
		}
	}

	return nil
}

func (e ExprRunner) hasOption(option string) bool {
	_, ok := e.optionsMap[option]
	return ok
}

func (e *ExprRunner) evalBaseExpression() error {
	// Base expression
	switch {
	case e.A && e.B && !e.C:
		e.H = M
	case e.A && e.B && e.C:
		e.H = P
	case !e.A && e.B && e.C:
		e.H = T
	default:
		return errors.New("no base expression validated to true")
	}

	switch e.H {
	case M:
		e.K = e.D + (e.D * float64(e.E) / 10)
	case P:
		e.K = e.D + (e.D * float64(e.E-e.F) / 25.5)
	case T:
		e.K = e.D - (e.D * float64(e.F) / 30)
	}

	return nil
}

func (e *ExprRunner) evalCustomExpression1() {
	if e.H == M {
		e.K = 2*e.D + (e.D * float64(e.E) / 100)
	}
}

func (e *ExprRunner) evalCustomExpression2() {
	if e.A && e.B && !e.C {
		e.H = T
	} else if e.A && !e.B && e.C {
		e.H = M
	}

	if e.H == M {
		e.K = float64(e.F) + e.D + (e.D * float64(e.E) / 100)
	}
}

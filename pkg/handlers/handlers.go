package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Expression struct {
	A bool    `json:"a"`
	B bool    `json:"b"`
	C bool    `json:"c"`
	D float64 `json:"d"`
	E int     `json:"e"`
	F int     `json:"f"`
	H float64
	K float64
}

const M, P, T = 1.0, 2.0, 3.0

func (e *Expression) EvalBaseExpression() error {
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

	if e.H == M {
		e.K = e.D + (e.D * float64(e.E) / 10)
	}

	if e.H == P {
		e.K = e.D + (e.D * float64(e.E-e.F) / 25.5)
	}

	if e.H == T {
		e.K = e.D - (e.D * float64(e.F) / 30)
	}

	return nil
}

func (e *Expression) EvalCustomExpression1() {
	if e.H == M {
		e.K = 2*e.D + (e.D * float64(e.E) / 100)
	}
}

func (e *Expression) EvalCustomExpression2() {
	if e.A && e.B && !e.C {
		e.H = T
	} else if e.A && !e.B && e.C {
		e.H = M
	} else if e.H == M {
		e.K = float64(e.F) + e.D + (e.D * float64(e.E) / 100)
	}
}

// GetExpressionJSON receives json as input, transforms it based on expressions and outputs the resuls as json
func GetExpressionJSON(w http.ResponseWriter, r *http.Request) {
	var expr Expression

	if r.Method != "GET" {
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&expr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid input json provided - %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = expr.EvalBaseExpression()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expr.EvalCustomExpression1()
	expr.EvalCustomExpression2()

	output := struct {
		H float64 `json:"h"`
		K float64 `json:"k"`
	}{
		H: expr.H,
		K: expr.K,
	}

	out, err := json.Marshal(output)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable for encode output json - %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

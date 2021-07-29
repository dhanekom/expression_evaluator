package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhanekom/expression_evaluator/pkg/expressions"
)

// GetExpressionJSON receives json as input, transforms it based on expressions and outputs the resuls as json
func GetExpressionJSON(w http.ResponseWriter, r *http.Request) {
	var expr expressions.ExprRunner

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

	err = expr.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

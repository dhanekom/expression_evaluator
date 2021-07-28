package expressions

import (
	"encoding/json"
	"strings"
	"testing"
)

var tests = []struct {
	Description    string
	Input          string
	ExpectedOutput string
	ExpectedResult bool
}{
	{"No expression options", `{"a":true,"b":true,"c":false,"d":1,"e":2,"f":3}`, `no expression options found`, false},
	{"Duplicate expression options", `{"a":false,"b":false,"c":true, "options": ["BASE", "BASE"]}`, `option "BASE" used multiple times`, false},
	{"Invalid expression options", `{"a":false,"b":false,"c":true, "options": ["NO_EXIST"]}`, `"NO_EXIST" is not a valid expression option`, false},
	{"No base expressions validated to true", `{"a":false,"b":false,"c":true, "options": ["BASE"]}`, `no base expression validated to true`, false},
	{"Success non-existent input json attributes ignored", `{"a":true,"b":true,"no-exist":true,"options": ["BASE"]}`, `{"h":1,"k":0}`, true},
	{"Success Options BASE expression 1", `{"a":true,"b":true,"c":false,"d":1,"e":2,"f":3,"options": ["BASE"]}`, `{"h":1,"k":1.2}`, true},
	{"Success Options BASE expression 2", `{"a":true,"b":true,"c":true,"d":1,"e":2,"f":3,"options": ["BASE"]}`, `{"h":2,"k":0.9607843137254902}`, true},
	{"Success Options BASE expression 3", `{"a":false,"b":true,"c":true,"d":1,"e":2,"f":3,"options": ["BASE"]}`, `{"h":3,"k":0.9}`, true},
	{"Success Options CUSTOM1 expression 1", `{"a":true,"b":true,"c":false,"d":1,"e":2,"f":3,"options": ["BASE", "CUSTOM1"]}`, `{"h":1,"k":2.02}`, true},
	{"Success Options CUSTOM2 expression 1", `{"a":true,"b":true,"c":false,"d":1,"e":2,"f":3,"options": ["BASE", "CUSTOM1", "CUSTOM2"]}`, `{"h":3,"k":2.02}`, true},
	{"Success Options CUSTOM2 expression 2", `{"a":true,"b":false,"c":true,"d":1,"e":2,"f":3,"options": ["CUSTOM2"]}`, `{"h":1,"k":4.02}`, true},
	{"Success Options CUSTOM2 expression 3", `{"a":true,"b":true,"c":false,"d":1,"e":2,"f":3,"options": ["CUSTOM2"]}`, `{"h":3,"k":0}`, true},
}

func TestRun(t *testing.T) {
	for _, e := range tests {
		t.Run(e.Description, func(t *testing.T) {
			var expr ExprRunner
			decoder := json.NewDecoder(strings.NewReader(e.Input))
			err := decoder.Decode(&expr)
			if err != nil {
				if e.ExpectedResult {
					t.Error(err.Error())
				}
			} else {
				err = expr.Run()
				if err != nil {
					if e.ExpectedResult {
						t.Error(err.Error())
					} else if e.ExpectedOutput != "" && e.ExpectedOutput != err.Error() {
						t.Errorf(`expected output "%s", got "%s"`, e.ExpectedOutput, err.Error())
					}
				} else {
					output := struct {
						H float64 `json:"h"`
						K float64 `json:"k"`
					}{
						H: expr.H,
						K: expr.K,
					}

					out, _ := json.Marshal(output)

					if !e.ExpectedResult {
						t.Errorf("expected failure but success - output was %s", string(out))
					} else if string(out) != e.ExpectedOutput {
						t.Errorf("expected output %s, got %s", e.ExpectedOutput, string(out))
					}
				}
			}
		})
	}
}

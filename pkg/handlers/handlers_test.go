package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var tests = []struct {
	Description        string
	Method             string
	Input              string
	ExpectedOutput     string
	ExpectedStatusCode int
}{
	{"Only GET methods allowed", "POST", "", "invalid http method", http.StatusMethodNotAllowed},
	{"Invalid JSON", "GET", `{`, "", http.StatusBadRequest},
	{"No base expressions validated to true", "GET", `{"a":false,"b":false,"c":true, "option": "BASE"}`, `no base expression validated to true`, http.StatusBadRequest},
	{"Success non-existent input json attributes ignored", "GET", `{"a":true,"b":true,"no-exist":true,"options": ["BASE"]}`, `{"h":1,"k":0}`, http.StatusOK},
	{"Success Options BASE expression 1", "GET", `{"a":true,"b":true,"c":false,"d":1,"e":2,"f":3,"options": ["BASE"]}`, `{"h":1,"k":1.2}`, http.StatusOK},
}

func TestRootHandler(t *testing.T) {
	for _, e := range tests {
		t.Run(e.Description, func(t *testing.T) {
			reader := strings.NewReader(e.Input)

			req := httptest.NewRequest(e.Method, "/some-url", reader)

			w := httptest.NewRecorder()
			GetExpressionJSON(w, req)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if e.ExpectedStatusCode != w.Code {
				t.Errorf(`want http status %d, got %d - %s`, e.ExpectedStatusCode, w.Code, string(body))
			} else if w.Code == http.StatusOK {
				if e.ExpectedOutput != string(body) {
					t.Errorf(`want output "%s", got "%s"`, e.ExpectedOutput, string(body))
				}
			}
		})
	}
}

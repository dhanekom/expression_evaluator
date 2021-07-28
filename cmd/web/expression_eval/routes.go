package main

import (
	"net/http"

	"github.com/dhanekom/wc_assessment/pkg/handlers"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/expression-json", handlers.GetExpressionJSON)

	return mux
}

package middleware

import (
	"net/http"
)

// Use applies a list of middleware functions to an http.Handler.
// Middlewares are applied in reverse order so that the first middleware
// wraps the handler last, preserving the expected chaining behavior.
func Use(handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) http.Handler {
	h := http.HandlerFunc(handler)
	for i := len(middlewares); i > 0; i-- {
		h = middlewares[i-1](h)
	}

	return h
}

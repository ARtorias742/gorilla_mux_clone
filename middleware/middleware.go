package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Simple logging example
		println("Request:", r.Method, r.URL.Path)
		next(w, r)
	}
}

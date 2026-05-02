package middleware

import (
	"net/http"
)

func Preflight(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle preflight (OPTIONS request)
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return

		} // Forward request to mux router
		next.ServeHTTP(w, r)

	})
}

package middleware

import (
	"net/http"
)

// Preflight → browser এর special request (OPTIONS) handle করে
// 👉 এটা mainly CORS এর part

func Preflight(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// ------------------ CHECK OPTIONS REQUEST ------------------
		// browser first এ OPTIONS request পাঠায় (preflight check)
		if r.Method == "OPTIONS" {

			// 200 OK return করলেই browser satisfied
			w.WriteHeader(200)
			return
		}

		// ------------------ NORMAL FLOW ------------------
		// যদি OPTIONS না হয় → next middleware / handler এ যাবে
		next.ServeHTTP(w, r)
	})
}

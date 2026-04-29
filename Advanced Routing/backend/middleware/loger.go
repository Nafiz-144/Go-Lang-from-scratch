package middleware

import (
	"log"
	"net/http"
	"time"
)

// Loger → request log করার জন্য
// 👉 debugging + monitoring এর জন্য খুব useful

func Loger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// ------------------ START TIME ------------------
		start := time.Now()

		// ------------------ CALL NEXT ------------------
		// handler execute হবে
		next.ServeHTTP(w, r)

		// ------------------ AFTER RESPONSE ------------------
		// request complete হওয়ার পর log print
		log.Println(
			r.Method,          // GET / POST
			r.URL.Path,        // /products
			time.Since(start), // কত সময় লাগলো
		)
	})
}

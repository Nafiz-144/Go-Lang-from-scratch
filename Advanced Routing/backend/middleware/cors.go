package middleware

import "net/http"

// Cors → browser (frontend) থেকে request allow করার জন্য ব্যবহার হয়
// 👉 Without this → React/Frontend request block হয়ে যাবে

func Cors(next http.Handler) http.Handler {

	// আমরা নতুন handler return করছি
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// ------------------ CORS HEADERS ------------------

		// সব origin (frontend) কে allow করছে
		// ⚠️ production এ "*" না দিয়ে specific domain use করা ভালো
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// কোন HTTP method allow
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")

		// কোন headers client send করতে পারবে
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Habib")

		// default response type JSON
		w.Header().Set("Content-Type", "application/json")

		// ------------------ CALL NEXT ------------------
		// 👉 next মানে পরের middleware / handler
		// এটা না ডাকলে request এখানেই শেষ হয়ে যাবে
		next.ServeHTTP(w, r)
	})
}

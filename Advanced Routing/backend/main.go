package main

import "nafiz/cmd"

// ====================== MAIN FUNCTION ======================

func main() {
	cmd.Serve()

}

// ====================== MIDDLEWARE ======================
/*
// CORS middleware (applied per route)
func crosMiddleware(next http.Handler) http.Handler {

	handleCors := func(w http.ResponseWriter, r *http.Request) {

		// Allow all origins (for development)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allowed HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTION")

		// Allowed headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Response type
		w.Header().Set("Content-Type", "application/json")

		// Call next handler
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handleCors)
}*/

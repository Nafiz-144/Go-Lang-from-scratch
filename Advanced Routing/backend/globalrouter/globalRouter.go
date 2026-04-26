package globalrouter

import "net/http"

// Global middleware (handles all incoming requests)
func GlobalRouter(mux *http.ServeMux) http.Handler {

	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Habib")
		w.Header().Set("Content-Type", "application/json")

		// Handle preflight (OPTIONS request)
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return

		} // Forward request to mux router
		mux.ServeHTTP(w, r)

	}

	return http.HandlerFunc(handleAllReq)
}

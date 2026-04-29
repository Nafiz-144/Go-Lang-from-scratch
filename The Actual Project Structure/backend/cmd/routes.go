package cmd

import (
	"nafiz/handlers"
	"nafiz/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Test route
	mux.Handle("GET /rahim", manager.With(
		http.HandlerFunc(handlers.Test),
	))

	// Get all products
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.Getproduct),
	))

	// Add product
	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.Addproduct),
	))

	// Get product by ID
	mux.Handle("GET /products/{productId}", manager.With(
		http.HandlerFunc(handlers.GetproductByID),
	))
}

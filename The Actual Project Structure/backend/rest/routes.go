package rest

import (
	"nafiz/rest/handlers"
	middleware "nafiz/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// Get all products
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.Getproducts),
	))

	// Add product
	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.Addproduct),
	))

	// Get product by ID
	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(handlers.Getproduct),
	))

	// Get product by ID
	mux.Handle("PUT /products/{id}", manager.With(
		http.HandlerFunc(handlers.Updateproduct),
	))

	mux.Handle("DELETE /products/{id}", manager.With(
		http.HandlerFunc(handlers.Deleteproduct),
	))

	mux.Handle("POST /users", manager.With(
		http.HandlerFunc(handlers.CreateUser),
	))

	mux.Handle("POST /users/login", manager.With(
		http.HandlerFunc(handlers.Login),
	))

}

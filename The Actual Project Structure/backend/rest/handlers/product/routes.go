package product

import (
	middleware "nafiz/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(h.Getproducts),
	))

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(h.Addproduct),
		h.middlewares.AuthenticateJWT,
	))

	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(h.Getproduct),
	))

	mux.Handle("PUT /products/{id}", manager.With(
		http.HandlerFunc(h.Updateproduct), h.middlewares.AuthenticateJWT,
	))

	mux.Handle("DELETE /products/{id}", manager.With(
		http.HandlerFunc(h.Deleteproduct), h.middlewares.AuthenticateJWT,
	))
}

package cmd

import (
	"fmt"
	"net/http"

	"nafiz/globalrouter"
	"nafiz/handler"
)

func Serve() {

	mux := http.NewServeMux()

	// Get all products
	mux.Handle("GET /products", http.HandlerFunc(handler.Getproduct))

	// Add product
	mux.Handle("POST /products", http.HandlerFunc(handler.Addproduct))

	// Get product by ID
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handler.GetproductByID))

	fmt.Println("Server Running on:8080")

	err := http.ListenAndServe(":8080", globalrouter.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error Starting the Server:", err)
	}
}

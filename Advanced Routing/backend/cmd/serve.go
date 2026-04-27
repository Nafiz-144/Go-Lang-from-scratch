package cmd

import (
	"fmt"
	"net/http"

	"nafiz/globalrouter"
	"nafiz/handlers"
	"nafiz/middleware"
)

func Serve() {

	mux := http.NewServeMux()

	mux.Handle("GET /nafiz", middleware.Hudai(middleware.Loger(http.HandlerFunc(handlers.Test))))

	// Get all products
	mux.Handle("GET /products", middleware.Hudai(middleware.Loger(http.HandlerFunc(handlers.Getproduct))))

	// Add product
	mux.Handle("POST /products", middleware.Hudai(middleware.Loger(http.HandlerFunc(handlers.Addproduct))))

	// Get product by ID
	mux.Handle("GET /products/{productId}", middleware.Hudai(middleware.Loger(http.HandlerFunc(handlers.GetproductByID))))

	fmt.Println("Server Running on:8080")

	err := http.ListenAndServe(":8080", globalrouter.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error Starting the Server:", err)
	}
}

package handlers

import (
	"encoding/json"
	"fmt"
	"nafiz/database"

	"nafiz/utill"
	"net/http"
)

// POST /addproduct
// Adds a new product to ProductList
func Addproduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	// Decode JSON request body into struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please enter a valid JSON", 400)
		return
	}

	// Auto-generate ID
	newProduct.ID = len(database.ProductList) + 1

	// Add product to slice
	database.ProductList = append(database.ProductList, newProduct)

	// Send response
	utill.SendData(w, newProduct, 201)
}

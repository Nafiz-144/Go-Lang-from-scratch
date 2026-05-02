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
		http.Error(w, "Please enter a valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	// Send response
	utill.SendData(w, createdProduct, 201)
}

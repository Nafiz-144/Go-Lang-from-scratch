package handlers

import (
	"encoding/json"
	"nafiz/database"
	"net/http"
)

// Addproduct → নতুন product add করে
// URL: POST /products

func Addproduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// ------------------ READ REQUEST BODY ------------------
	var newProduct database.Product

	// client থেকে আসা JSON → Go struct এ convert
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// ------------------ SIMPLE ID GENERATION ------------------
	// last product এর ID + 1
	if len(database.ProductList) > 0 {
		newProduct.ID = database.ProductList[len(database.ProductList)-1].ID + 1
	} else {
		newProduct.ID = 1
	}

	// ------------------ ADD TO DATABASE ------------------
	database.ProductList = append(database.ProductList, newProduct)

	// ------------------ RESPONSE ------------------
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(newProduct)
}

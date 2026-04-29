package handlers

import (
	"encoding/json"
	"nafiz/database"
	"net/http"
	"strconv"
)

// GetproductByID → নির্দিষ্ট product return করে
// URL: /products/{productId}

func GetproductByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// ------------------ GET ID FROM URL ------------------
	// URL থেকে productId বের করা
	idParam := r.PathValue("productId")

	// string → int convert
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// ------------------ SEARCH PRODUCT ------------------
	for _, product := range database.ProductList {

		// যদি ID match করে
		if product.ID == id {

			// product return
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	// ------------------ NOT FOUND ------------------
	http.Error(w, "Product not found", http.StatusNotFound)
}

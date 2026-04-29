package handlers

import (
	"encoding/json"
	"nafiz/database"
	"net/http"
)

// Getproduct → সব product client কে return করে
// যখন client GET /products call করে তখন এটা run হয়

func Getproduct(w http.ResponseWriter, r *http.Request) {

	// ------------------ RESPONSE HEADER ------------------
	// browser/client কে বলছি আমরা JSON পাঠাচ্ছি
	w.Header().Set("Content-Type", "application/json")

	// ------------------ ENCODE DATA ------------------
	// Go struct (ProductList) → JSON এ convert করে client কে পাঠানো
	err := json.NewEncoder(w).Encode(database.ProductList)

	if err != nil {
		// যদি encode fail করে
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}
}

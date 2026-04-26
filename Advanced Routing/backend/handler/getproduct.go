package handler

import (
	"nafiz/database"

	"nafiz/utill"
	"net/http"
)

// GET /getproduct
// Returns all products as JSON
func Getproduct(w http.ResponseWriter, r *http.Request) {
	utill.SendData(w, database.ProductList, 200)
}

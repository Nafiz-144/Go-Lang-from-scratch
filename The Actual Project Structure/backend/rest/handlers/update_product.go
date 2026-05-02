package handlers

import (
	"encoding/json"
	"fmt"
	"nafiz/database"
	"nafiz/utill"
	"net/http"
	"strconv"
)

// GET /getproduct
// Returns all products as JSON
func Updateproduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid Product id", 400)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please enter a valid JSON", 400)
		return
	}
	newProduct.ID = pId
	database.Update(newProduct)
	utill.SendData(w, "Successfully Update Product", 201)

}

// `struct
// pId
// jo conver`

package product

import (
	"encoding/json"
	"fmt"
	"nafiz/database"

	"nafiz/utill"
	"net/http"
)

func (h *Handler) Addproduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please enter a valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	utill.SendData(w, createdProduct, 201)
}

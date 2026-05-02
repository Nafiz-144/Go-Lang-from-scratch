package handlers

import (
	"nafiz/database"
	"nafiz/utill"
	"net/http"
	"strconv"
)

func Getproduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid Product id", 400)
		return
	}
	product := database.Get(pId)
	if product == nil {
		utill.SendError(w, 404, "Product not found!")
		return
	}

	utill.SendData(w, product, 400)

}

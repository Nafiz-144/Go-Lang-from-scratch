package handler

import (
	"nafiz/database"
	"nafiz/utill"
	"net/http"
	"strconv"
)

func GetproductByID(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("productId")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid Product id", 400)
		return
	}
	for _, product := range database.ProductList {
		if product.ID == pId {
			utill.SendData(w, product, 200)
			return
		}
	}
	utill.SendData(w, "Data not found", 400)

}

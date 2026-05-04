package product

import (
	"nafiz/database"
	"nafiz/utill"
	"net/http"
	"strconv"
)

// GET /getproduct
// Returns all products as JSON
func (h *Handler) Deleteproduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid Product id", 400)
		return
	}

	database.Delete(pId)
	utill.SendData(w, "Successfully Deleted Product", 201)

}

// `struct
// pId
// jo conver`

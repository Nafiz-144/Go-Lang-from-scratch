package product

import (
	"nafiz/database"

	"nafiz/utill"
	"net/http"
)

func (h *Handler) Getproducts(w http.ResponseWriter, r *http.Request) {
	utill.SendData(w, database.List(), 200)
}

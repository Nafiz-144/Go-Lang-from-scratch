package product

import (
	"nafiz/utill"
	"net/http"
)

func (h *Handler) Getproducts(w http.ResponseWriter, r *http.Request) {

	productList, err := h.svc.List()
	if err != nil {
		utill.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utill.SendData(w, http.StatusOK, productList)
}

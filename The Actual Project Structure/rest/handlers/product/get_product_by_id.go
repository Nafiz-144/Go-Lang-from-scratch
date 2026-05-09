package product

import (
	"nafiz/utill"
	"net/http"
	"strconv"
)

func (h *Handler) Getproduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		utill.SendError(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	product, err := h.svc.Get(pId)
	if err != nil {
		utill.SendError(w, http.StatusInternalServerError, "Internal server Error..")

		return

	}

	if product == nil {
		utill.SendError(w, http.StatusNotFound, "Product not found!")
		return
	}

	utill.SendData(w, http.StatusOK, product)

}

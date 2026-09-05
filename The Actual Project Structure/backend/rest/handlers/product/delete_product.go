package product

import (
	"fmt"
	"nafiz/utill"
	"net/http"
	"strconv"
)

func (h *Handler) Deleteproduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		utill.SendError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	err = h.productRepo.Delete(pId)
	if err != nil {
		fmt.Println("Error:", err)
		utill.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return

	}

	utill.SendData(w, http.StatusOK, "Successfully Deleted Product")

}

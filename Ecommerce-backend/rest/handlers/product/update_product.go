package product

import (
	"encoding/json"
	"fmt"
	"nafiz/domain"
	"nafiz/utill"
	"net/http"
	"strconv"
)

type ReqUpdateproduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

// GET /getproduct
// Returns all products as JSON
func (h *Handler) Updateproduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		utill.SendError(w, http.StatusBadRequest, "Inavalid product ID ")

		return
	}

	var req ReqUpdateproduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utill.SendError(w, http.StatusBadRequest, "Invalid Request Body")

		return
	}
	_, err = h.svc.Update(domain.Product{
		ID:          pId,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {

		utill.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utill.SendData(w, http.StatusOK, "Successfully Update Product")

}

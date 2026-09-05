package product

import (
	"encoding/json"
	"fmt"

	"nafiz/repo"
	"nafiz/utill"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func (h *Handler) Addproduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please enter a valid json", 400)
		utill.SendError(w, http.StatusBadRequest, "invalid req Body")
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{

		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})

	if err != nil {

		utill.SendError(w, http.StatusInternalServerError, "Internal Sever Error")
		return
	}
	utill.SendData(w, http.StatusCreated, createdProduct)
}

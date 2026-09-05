package user

import (
	"encoding/json"
	"fmt"

	"nafiz/utill"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var req ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utill.SendError(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	usr, err := h.userRepo.Find(req.Email, req.Password)

	if err != nil {
		utill.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, err := utill.CreateJwt(h.cnf.JwtSecretKey, utill.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utill.SendData(w, http.StatusCreated, accessToken)
}

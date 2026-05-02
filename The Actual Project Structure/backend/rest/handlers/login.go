package handlers

import (
	"encoding/json"
	"fmt"

	"nafiz/database"
	"nafiz/utill"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var reqLogin ReqLogin

	// Decode JSON request body into struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}
	usr := database.Find(reqLogin.Email, reqLogin.Password)

	if usr == nil {
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
		return
	}

	utill.SendData(w, usr, http.StatusCreated)
}

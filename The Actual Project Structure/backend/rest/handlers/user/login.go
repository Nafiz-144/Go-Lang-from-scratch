package user

import (
	"encoding/json"
	"fmt"

	"nafiz/config"
	"nafiz/database"
	"nafiz/utill"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (*Handler) Login(w http.ResponseWriter, r *http.Request) {

	var reqLogin ReqLogin

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

	cnf := config.GetConfig()
	accessToken, err := utill.CreateJwt(cnf.JwtSecretKey, utill.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utill.SendData(w, accessToken, http.StatusCreated)
}

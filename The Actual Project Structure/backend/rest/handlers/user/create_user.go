package user

import (
	"encoding/json"
	"fmt"
	"nafiz/database"

	"nafiz/utill"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser database.User

	// Decode JSON request body into struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()

	// Send response
	utill.SendData(w, createdUser, http.StatusCreated)
}

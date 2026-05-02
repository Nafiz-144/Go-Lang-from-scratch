package handlers

import (
	"nafiz/database"

	"nafiz/utill"
	"net/http"
)

func Getproducts(w http.ResponseWriter, r *http.Request) {
	utill.SendData(w, database.List(), 200)
}

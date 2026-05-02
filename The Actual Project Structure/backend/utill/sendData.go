package utill

import (
	"encoding/json"
	"net/http"
)

// ====================== HELPER FUNCTION ======================

// Sends JSON response with status code
func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
func SendError(w http.ResponseWriter, statusCode int, msg string) {

	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}

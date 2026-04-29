package utill

import (
	"encoding/json"
	"net/http"
)

// ====================== HELPER FUNCTION ======================

// SendData → client কে JSON response পাঠানোর জন্য reusable function
// 👉 handler এ বারবার json encode + status code লেখার ঝামেলা কমায়

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {

	// ------------------ STATUS CODE ------------------
	// client কে জানাচ্ছি request success / error type
	// যেমন: 200, 201, 404, 500
	w.WriteHeader(statusCode)

	// ------------------ CREATE JSON ENCODER ------------------
	// encoder তৈরি করছি → Go data → JSON এ convert করবে
	encoder := json.NewEncoder(w)

	// ------------------ SEND RESPONSE ------------------
	// data → JSON → client এ পাঠানো হচ্ছে
	encoder.Encode(data)
}

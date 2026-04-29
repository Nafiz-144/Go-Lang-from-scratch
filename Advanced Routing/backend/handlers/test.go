package handlers

import (
	"log"
	"net/http"
)

// Test → এটা একটা HTTP handler function
// 👉 যখন এই route hit হবে (e.g. GET /rahim), তখন এই function execute হবে

func Test(w http.ResponseWriter, r *http.Request) {

	// ------------------ LOGGING ------------------
	// server console এ print করবে
	// 👉 client কিছুই দেখবে না (important!)
	log.Println("I am Handeler")

	// ⚠️ NOTE:
	// এখানে client কে কোনো response পাঠানো হয়নি
	// তাই browser/Postman request করলে empty response পাবে
}

package main

import (
	"fmt"
	"net/http"
)

func halloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Sadman. I am a software engineer")
}
func main() {
	mux := http.NewServeMux()              //router
	mux.HandleFunc("/hello", halloHandler) // route
	mux.HandleFunc("/about", aboutHandler) //route
	fmt.Println("Server Runing on:3000")
	err := http.ListenAndServe(":3000", mux) //star a server in a specific port
	if err != nil {
		fmt.Println("Error Starting the Server:", err)
	}
}

// Client → Network → NIC → Kernel → Socket → Go runtime → Goroutine → Handler → Response → Kernel → NIC → Client

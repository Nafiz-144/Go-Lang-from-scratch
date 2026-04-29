package cmd

import (
	"fmt"
	"net/http"

	"nafiz/middleware"
)

func Serve() {

	manager := middleware.NewManager()

	mux := http.NewServeMux()

	// warppedMux := manager.WrapMux(mux, middleware.Loger,
	// 	middleware.Hudai, middleware.CorsWithPreflight)

	manager.Use(middleware.Preflight, middleware.Cors, middleware.Loger)

	WrapMux := manager.WrapMux(mux)
	initRoutes(mux, manager)
	fmt.Println("Server Running on:8080")

	err := http.ListenAndServe(":8080", WrapMux)
	if err != nil {
		fmt.Println("Error Starting the Server:", err)
	}
}

/*
->global router->hudai->loger->test

*/

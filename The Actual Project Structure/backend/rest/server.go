package rest

import (
	"fmt"
	"nafiz/config"
	middleware "nafiz/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {

	manager := middleware.NewManager()

	mux := http.NewServeMux()

	manager.Use(middleware.Preflight, middleware.Cors, middleware.Loger)

	WrapMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server Running on port:", addr)
	err := http.ListenAndServe(addr, WrapMux)
	if err != nil {
		fmt.Println("Error Starting the Server:", err)
		os.Exit(1)
	}
}

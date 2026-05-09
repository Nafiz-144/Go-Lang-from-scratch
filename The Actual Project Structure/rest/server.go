package rest

import (
	"fmt"
	"nafiz/config"
	"nafiz/rest/handlers/product"
	"nafiz/rest/handlers/user"
	middleware "nafiz/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}

}

func (server *Server) Start() {

	manager := middleware.NewManager()

	manager.Use(middleware.Preflight, middleware.Cors,
		middleware.Loger)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server Running on port:", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error Starting the Server:", err)
		os.Exit(1)
	}
}

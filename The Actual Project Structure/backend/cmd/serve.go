package cmd

import (
	"nafiz/config"
	"nafiz/rest"
	"nafiz/rest/handlers/product"
	"nafiz/rest/handlers/user"
	middleware "nafiz/rest/middlewares"
)

func Serve() {

	cnf := config.GetConfig()
	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()

}

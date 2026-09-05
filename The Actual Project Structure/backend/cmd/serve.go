package cmd

import (
	"fmt"
	"nafiz/config"
	"nafiz/infra/db"
	"nafiz/repo"
	"nafiz/rest"
	"nafiz/rest/handlers/product"
	"nafiz/rest/handlers/user"
	middleware "nafiz/rest/middlewares"
	"os"
)

func Serve() {

	cnf := config.GetConfig()
	dbCon, err := db.NewConnection()
	if err != nil {

		fmt.Println(err)
		os.Exit(1)

	}

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(dbCon)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()

}

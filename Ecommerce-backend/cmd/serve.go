package cmd

import (
	"fmt"
	"nafiz/config"
	"nafiz/infra/db"
	"nafiz/product"
	"nafiz/repo"
	"nafiz/rest"
	prdcthandler "nafiz/rest/handlers/product"
	usrHandler "nafiz/rest/handlers/user"
	middleware "nafiz/rest/middlewares"
	"nafiz/user"
	"os"
)

func Serve() {

	cnf := config.GetConfig()
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {

		fmt.Println(err)
		os.Exit(1)

	}
	err = db.MigrateDB(dbCon, "./migrations")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)
	//domains
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := prdcthandler.NewHandler(middlewares, prdctSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()

}

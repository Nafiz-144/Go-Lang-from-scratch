package cmd

import (
	"nafiz/config"
	"nafiz/rest"
)

func Serve() {

	rest.Start(config.GetConfig())

}

/*
->global router->hudai->loger->test

*/

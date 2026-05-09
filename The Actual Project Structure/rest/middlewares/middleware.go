package middleware

import "nafiz/config"

type Misddlewares struct {
	cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Misddlewares {
	return &Misddlewares{
		cnf: cnf,
	}
}

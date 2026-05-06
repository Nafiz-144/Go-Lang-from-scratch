package product

import (
	"nafiz/repo"
	middleware "nafiz/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Misddlewares
	productRepo repo.ProductRepo
}

func NewHandler(middlewares *middleware.Misddlewares,
	productRepo repo.ProductRepo,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}

}

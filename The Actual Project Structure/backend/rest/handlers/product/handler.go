package product

import middleware "nafiz/rest/middlewares"

type Handler struct {
	middlewares *middleware.Misddlewares
}

func NewHandler(middlewares *middleware.Misddlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}

}

package product

import (
	middleware "nafiz/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Misddlewares
	svc         Service
}

func NewHandler(middlewares *middleware.Misddlewares,
	svc Service,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc:         svc,
	}

}

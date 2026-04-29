package middleware

import (
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	globalMiddleware []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middleware ...Middleware) {
	mngr.globalMiddleware = append(mngr.globalMiddleware, middleware...)

}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {

	h := handler
	// //middleware =[hudai,logger] inx=hudai->0...
	// for i := len(middlewares) - 1; i >= 0; i-- {
	// 	middleware := middlewares[i] //loger
	// 	n = middleware(n)            //   middleware.Loger(http.HandlerFunc(handlers.Getproduct))
	// }
	//n=middleware.Loger)((middleware.Hudai, middleware.Loger)(http.HandlerFunc(handlers.Getproduct)))
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	//global middle ware
	// for _, globalMiddleware := range mngr.globalMiddleware {
	// 	h = globalMiddleware(h)
	// }

	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler

	for _, middleware := range mngr.globalMiddleware {
		h = middleware(h)
	}

	return h
}

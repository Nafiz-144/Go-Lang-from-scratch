package middleware

import (
	"net/http"
)

// ====================== MIDDLEWARE TYPE ======================

// Middleware → function type
// 👉 এটা এমন function যা একটা handler নেয়
// 👉 এবং নতুন handler return করে (wrapped version)

// Think:
// oldHandler → middleware → newHandler

type Middleware func(next http.Handler) http.Handler

// ====================== MANAGER STRUCT ======================

// Manager → middleware manage করার জন্য
// 👉 এখানে global middleware store করা হয়

type Manager struct {
	globalMiddleware []Middleware
}

// ====================== CONSTRUCTOR ======================

// NewManager → নতুন Manager create করে
// 👉 শুরুতে empty middleware list থাকে

func NewManager() *Manager {
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}
}

// ====================== ADD GLOBAL MIDDLEWARE ======================

// Use → global middleware add করার জন্য
// 👉 এই middleware সব route এ apply হবে

func (mngr *Manager) Use(middleware ...Middleware) {

	// append → list এ add করে
	mngr.globalMiddleware = append(mngr.globalMiddleware, middleware...)
}

// ====================== ROUTE-SPECIFIC MIDDLEWARE ======================

// With → specific handler এর সাথে middleware apply করে
// 👉 only this route এর জন্য middleware use করা হয়

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {

	h := handler

	// ------------------ APPLY MIDDLEWARE ------------------

	// ⚠️ Current logic (forward loop)
	for _, middleware := range middlewares {

		// handler → wrap হচ্ছে
		h = middleware(h)
	}

	// 👉 final wrapped handler return
	return h
}

// ====================== APPLY GLOBAL MIDDLEWARE ======================

// WrapMux → পুরো router কে middleware দিয়ে wrap করে
// 👉 সব request আগে এখানে আসবে

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler

	// global middleware apply হচ্ছে
	for _, middleware := range mngr.globalMiddleware {
		h = middleware(h)
	}

	return h
}

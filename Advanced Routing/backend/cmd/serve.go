package cmd

import (
	"fmt"
	"net/http"

	"nafiz/middleware"
)

// Serve() হচ্ছে তোমার server start করার main function
// 👉 এখানে server setup + middleware setup + routing setup হয়

func Serve() {

	// ------------------ CREATE MIDDLEWARE MANAGER ------------------
	// এটা middleware handle করবে
	manager := middleware.NewManager()

	// ------------------ CREATE ROUTER ------------------
	// ServeMux → built-in router
	mux := http.NewServeMux()

	// ------------------ REGISTER GLOBAL MIDDLEWARE ------------------
	// manager.Use(...) → global middleware add করা হয়
	// এগুলো সব route এ apply হবে

	manager.Use(
		middleware.Preflight, // OPTIONS request handle (important for browser)
		middleware.Cors,      // CORS allow করে (frontend থেকে request আসার জন্য)
		middleware.Loger,     // logging (request info print করবে)
	)

	// ------------------ WRAP ROUTER WITH MIDDLEWARE ------------------
	// WrapMux → পুরো mux কে middleware দিয়ে wrap করে
	WrapMux := manager.WrapMux(mux)

	// ------------------ INIT ROUTES ------------------
	// সব route register করা
	initRoutes(mux, manager)

	fmt.Println("Server Running on:8080")

	// ------------------ START SERVER ------------------
	// ":8080" → port
	err := http.ListenAndServe(":8080", WrapMux)

	if err != nil {
		fmt.Println("Error Starting the Server:", err)
	}
}

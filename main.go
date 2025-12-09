package main

import (
	"authvault/handlers"
	"authvault/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("AuthVault running at http.localhost:8080")

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/refresh", handlers.RefreshToken)
	http.HandleFunc("/profile", middleware.AuthMiddleware(handlers.Profile))
	http.HandleFunc("/logout", middleware.AuthMiddleware(handlers.Logout))
	http.HandleFunc("/admin", middleware.AuthMiddleware(middleware.AdminOnly(handlers.AdminRoute)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"authsession/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/register", handlers.RegisterPage)
	http.HandleFunc("/register-submit", handlers.Register)
	http.HandleFunc("/login", handlers.LoginPage)
	http.HandleFunc("/login-submit", handlers.Login)
	http.HandleFunc("/welcome", handlers.WelcomePage)
	http.HandleFunc("/logout", handlers.Logout)

	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

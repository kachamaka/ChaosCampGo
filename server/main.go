package main

import (
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/handlers"
)

func main() {
	database.Get().Connect()
	defer database.Get().Disconnect()
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	// http.HandleFunc("/getUsers", middleware.Auth(getUsers))
	// http.HandleFunc("/getEvents", hello)
	// http.HandleFunc("/addEvent", hello)
	// http.HandleFunc("/deleteEvent", hello)

	// tokens.GenerateToken("id")
	http.ListenAndServe(":8888", nil)
}

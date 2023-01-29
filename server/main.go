package main

import (
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/handlers"
)

func main() {
	database.Get().Connect()
	defer database.Get().Disconnect()
	// http.HandleFunc("/login", headers)
	http.HandleFunc("/register", handlers.RegisterHandler)
	// http.HandleFunc("/getEvents", hello)
	// http.HandleFunc("/addEvent", hello)
	// http.HandleFunc("/deleteEvent", hello)

	http.ListenAndServe(":8888", nil)
}

package main

import (
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/handlers"
	"github.com/kachamaka/chaosgo/middleware"
)

func main() {
	database.Get().Connect()
	defer database.Get().Disconnect()

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/addEvent", middleware.Auth(handlers.AddEventHandler))
	http.HandleFunc("/deleteEvent", middleware.Auth(handlers.DeleteEventHandler))
	http.HandleFunc("/getEvents", middleware.Auth(handlers.GetEventsHandler))

	http.ListenAndServe(":8888", nil)
	// http.ListenAndServe(":8888", middleware.CORS(r))
}

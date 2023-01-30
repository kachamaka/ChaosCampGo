package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/handlers"
	"github.com/kachamaka/chaosgo/middleware"
	"github.com/kachamaka/chaosgo/models"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	ID := r.Context().Value("_id")
	fmt.Println(ID)

	_ = decoder

	encoder.Encode(models.BasicResponse{Success: true, Message: "all good"})
}

func main() {
	database.Get().Connect()
	defer database.Get().Disconnect()
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/getUsers", middleware.Auth(getUsers))
	// http.HandleFunc("/getEvents", hello)
	// http.HandleFunc("/addEvent", hello)
	// http.HandleFunc("/deleteEvent", hello)

	// tokens.GenerateToken("id")
	http.ListenAndServe(":8888", nil)
}

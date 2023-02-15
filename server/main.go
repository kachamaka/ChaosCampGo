package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kachamaka/chaosgo/config"
	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/handlers"
	"github.com/kachamaka/chaosgo/middleware"
)

var logger *string

func init() {
	logger = flag.String("log", "", "path to logger")
}

func main() {
	// Load config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("load config error:", err)
	}

	flag.Parse()

	// Configure logger
	if *logger != "" {
		file, err := os.Create("logs.txt")
		if err != nil {
			log.Fatal("error creating log file:", err)
		}
		log.SetOutput(file)
		fmt.Println("Logging in logs.txt")
		defer file.Close()
	} else {
		log.SetOutput(io.Discard)
	}

	database.Get().Connect(config.DatabaseAddress, config.DatabaseName)
	defer database.Get().Disconnect()

	// Send reminders
	go database.Get().SendReminders()

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/addEvent", middleware.Auth(handlers.AddEventHandler))
	mux.HandleFunc("/deleteEvent", middleware.Auth(handlers.DeleteEventHandler))
	mux.HandleFunc("/getEvents", middleware.Auth(handlers.GetEventsHandler))
	mux.HandleFunc("/addReminder", middleware.Auth(handlers.AddReminderHandler))

	fmt.Println("Running on:", config.ServerAddress)
	http.ListenAndServe(config.ServerAddress, middleware.CORS(mux))
}

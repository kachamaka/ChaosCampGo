package main

import (
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/handlers"
	"github.com/kachamaka/chaosgo/middleware"
	"github.com/kachamaka/chaosgo/models"
)

func main() {
	database.Get().Connect()
	defer database.Get().Disconnect()

	//send reminders
	go models.SendReminders()

	// reminder := models.Reminder{
	// 	UserID:    "63d8cdd577f897d88c753fbf",
	// 	Email:     "martin.popov42@gmail.com",
	// 	Subject:   "Work",
	// 	Time:      0,
	// 	StartTime: 0,
	// }
	// start := time.Now().Add(time.Minute * 30)
	// reminder.StartTime = start.Unix()
	// reminder.Time = start.Unix() - 900
	// reminder.Send()
	// return

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/addEvent", middleware.Auth(handlers.AddEventHandler))
	http.HandleFunc("/deleteEvent", middleware.Auth(handlers.DeleteEventHandler))
	http.HandleFunc("/getEvents", middleware.Auth(handlers.GetEventsHandler))
	http.HandleFunc("/addReminder", middleware.Auth(handlers.AddReminderHandler))

	http.ListenAndServe(database.Get().Config.ServerAddress, nil)
	// http.ListenAndServe(":8888", middleware.CORS(r))
}

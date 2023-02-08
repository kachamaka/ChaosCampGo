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

	//send reminders
	go database.SendReminders()

	//MAKE SENDING WITH MAILGUN

	// reminder := models.Reminder{
	// 	UserID:    "63d8cdd577f897d88c753fbf",
	// 	Email:     "martin.popov42@gmail.com",
	// 	Subject:   "Work",
	// 	Time:      0,
	// 	StartTime: 0,
	// }
	// start := time.Now().Add(time.Hour).Add(time.Second * 15)
	// reminder.StartTime = start.Unix()
	// reminder.Time = start.Unix() - 3600
	// database.Send(reminder)

	// Your available domain names can be found here:
	// (https://app.mailgun.com/app/domains)

	// return

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/addEvent", middleware.Auth(handlers.AddEventHandler))
	http.HandleFunc("/deleteEvent", middleware.Auth(handlers.DeleteEventHandler))
	http.HandleFunc("/getEvents", middleware.Auth(handlers.GetEventsHandler))
	http.HandleFunc("/addReminder", middleware.Auth(handlers.AddReminderHandler))
	// http.HandleFunc("/deleteReminder", middleware.Auth(handlers.DeleteReminderHandler))
	// http.HandleFunc("/deleteReminder", handlers.DeleteReminderHandler)

	http.ListenAndServe(database.Get().Config.ServerAddress, nil)
	// http.ListenAndServe(":8888", middleware.CORS(r))
}

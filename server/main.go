package main

import (
	"fmt"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/handlers"
	"github.com/kachamaka/chaosgo/middleware"
)

func main() {
	// from := mail.NewEmail("golangcc", "golangcc42@gmail.com")
	// to := mail.NewEmail("", "martin.popov42@gmail.com")
	// subject := fmt.Sprintf("Reminder for event: %s", "OK")
	// plainTextContent := "test"
	// htmlContent := plainTextContent
	// message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	// message.SendAt = 1676210160

	// client := sendgrid.NewSendClient("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")
	// _, err := client.Send(message)
	// if err != nil {
	// 	log.Println("sendgrid error: ", err)
	// 	return
	// }

	// return

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
	database.Get().Connect()
	defer database.Get().Disconnect()

	//send reminders
	go database.SendReminders()

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/addEvent", middleware.Auth(handlers.AddEventHandler))
	mux.HandleFunc("/deleteEvent", middleware.Auth(handlers.DeleteEventHandler))
	mux.HandleFunc("/getEvents", middleware.Auth(handlers.GetEventsHandler))
	mux.HandleFunc("/addReminder", middleware.Auth(handlers.AddReminderHandler))
	// http.HandleFunc("/deleteReminder", middleware.Auth(handlers.DeleteReminderHandler))
	// http.HandleFunc("/deleteReminder", handlers.DeleteReminderHandler)

	// http.ListenAndServe(database.Get().Config.ServerAddress, nil)
	fmt.Println("Running on:", database.Get().Config.ServerAddress)
	http.ListenAndServe(database.Get().Config.ServerAddress, middleware.CORS(mux))
}

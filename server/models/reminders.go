package models

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/kachamaka/chaosgo/database"
	"gopkg.in/mgo.v2/bson"
)

type Reminder struct {
	UserID    string `json:"user_id" bson:"user_id"`
	Email     string `json:"email" bson:"email"`
	Subject   string `json:"subject" bson:"subject"`
	Time      int64  `json:"time" bson:"time"`
	StartTime int64  `json:"startTime" bson:"startTime"`
}

type ReminderRequest struct {
	Subject   string `json:"subject"`
	TimeAhead int64  `json:"timeAhead"` //how many minutes earlier to send a reminder
	StartTime int64  `json:"startTime"`
}

func secondsToString(seconds int64) string {
	hours := seconds / 3600
	seconds -= hours * 3600
	minutes := seconds / 60
	seconds -= minutes * 60
	return fmt.Sprintf("%d hour(s), %d minute(s) and %d second(s)", hours, minutes, seconds)
}

func (r *Reminder) Send() {
	from := "golangcc42@gmail.com"
	pass := "gxrrzdnxobuxkdkj"
	to := r.Email
	body := fmt.Sprintf("Hello, your event \"%s\" is about to start in %s.", r.Subject, secondsToString(r.StartTime-r.Time))
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: Reminder for event %s\n\n%s", from, to, r.Subject, body)
	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")

	difference := r.Time - time.Now().Unix()
	// fmt.Println("SLEEP FOR:", time.Duration(difference*int64(time.Second)))
	time.Sleep(time.Duration(difference * int64(time.Second)))

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	reminders := database.Get().GetCollection(database.REMINDERS_COLLECTION)
	_, err = reminders.DeleteOne(context.TODO(), r)
	if err != nil {
		log.Println("error with deleting reminder after sending: ", err)
		return
	}

	log.Print("email sent")
}

func SendReminders() {
	remindersCollection := database.Get().GetCollection(database.REMINDERS_COLLECTION)
	cursor, err := remindersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("error fetching reminders: ", err)
		return
	}

	var reminders []Reminder
	err = cursor.All(context.TODO(), &reminders)
	if err != nil {
		log.Println("error decoding result: ", err)
		return
	}
	for _, reminder := range reminders {
		if reminder.Time < time.Now().Unix() {
			//reminder already due
			go func(reminder Reminder) {
				err := deleteReminder(reminder)
				if err != nil {
					log.Println("error deleting reminder: ", err)
				}
			}(reminder)
		} else {
			go reminder.Send()
		}
	}
}

func deleteReminder(reminder Reminder) error {
	reminders := database.Get().GetCollection(database.REMINDERS_COLLECTION)
	filter := bson.M{"user_id": reminder.UserID, "subject": reminder.Subject, "startTime": reminder.StartTime, "time": reminder.Time}

	_, err := reminders.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("error deleting reminder:", err)
		return err
	}

	log.Println("reminder deleted successfully")
	return nil
}

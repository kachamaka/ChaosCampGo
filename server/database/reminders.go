package database

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/kachamaka/chaosgo/models"
	"gopkg.in/mgo.v2/bson"
)

func secondsToString(seconds int64) string {
	hours := seconds / 3600
	seconds -= hours * 3600
	minutes := seconds / 60
	seconds -= minutes * 60
	return fmt.Sprintf("%d hour(s), %d minute(s) and %d second(s)", hours, minutes, seconds)
}

func Send(r models.Reminder) {
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

	reminders := Get().GetCollection(REMINDERS_COLLECTION)
	_, err = reminders.DeleteOne(context.TODO(), r)
	if err != nil {
		log.Println("error with deleting reminder after sending: ", err)
		return
	}

	log.Print("email sent")
}

func SendReminders() {
	remindersCollection := Get().GetCollection(REMINDERS_COLLECTION)
	cursor, err := remindersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("error fetching reminders: ", err)
		return
	}

	var reminders []models.Reminder
	err = cursor.All(context.TODO(), &reminders)
	if err != nil {
		log.Println("error decoding result: ", err)
		return
	}
	for _, reminder := range reminders {
		if reminder.Time < time.Now().Unix() {
			//reminder already due
			go func(reminder models.Reminder) {
				err := DeleteReminder(reminder)
				if err != nil {
					log.Println("error deleting reminder: ", err)
				}
			}(reminder)
		} else {
			go Send(reminder)
		}
	}
}

func DeleteReminder(reminder models.Reminder) error {
	reminders := Get().GetCollection(REMINDERS_COLLECTION)

	_, err := reminders.DeleteOne(context.TODO(), reminder)
	if err != nil {
		log.Println("error deleting reminder:", err)
		return err
	}

	log.Println("reminder deleted successfully")
	return nil
}

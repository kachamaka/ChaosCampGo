package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kachamaka/chaosgo/models"
	"github.com/mailgun/mailgun-go/v4"
	"gopkg.in/mgo.v2/bson"
)

func secondsToString(seconds int64) string {
	hours := seconds / 3600
	seconds -= hours * 3600
	minutes := seconds / 60
	seconds -= minutes * 60
	return fmt.Sprintf("%d hour(s), %d minute(s) and %d second(s)", hours, minutes, seconds)
}

func (db *Database) AddReminder(reminder models.Reminder) error {
	reminders := db.GetCollection(REMINDERS_COLLECTION)
	_, err := reminders.InsertOne(context.TODO(), reminder)
	if err != nil {
		log.Println("err adding reminder: ", err)
		//CUSTOM ERRORS
		return fmt.Errorf("error adding reminder")
	}

	return nil
}

func Send(r models.Reminder) {
	mg := mailgun.NewMailgun(MAILGUN_DOMAIN, MAILGUN_PRIVATE_API_KEY)

	sender := "golangcc42@gmail.com"
	subject := fmt.Sprintf("Reminder for event: %s", r.Subject)
	body := fmt.Sprintf("Hello, your event \"%s\" is about to start in %s.", r.Subject, secondsToString(r.StartTime-r.Time))
	recipient := r.Email

	message := mg.NewMessage(sender, subject, body, recipient)

	message.SetDeliveryTime(time.Unix(r.Time, 0))

	_, _, err := mg.Send(context.TODO(), message)

	if err != nil {
		log.Println("mailgun error: ", err)
		return
	}

	err = DeleteReminder(r)
	if err != nil {
		log.Println("error with deleting reminder after sending: ", err)
		return
	}

	log.Println("email sent")
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

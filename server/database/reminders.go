package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kachamaka/chaosgo/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// secondsToString converts seconds that is a unix timestamp to readable text for humans
func secondsToString(seconds int64) string {
	hours := seconds / 3600
	seconds -= hours * 3600
	minutes := seconds / 60
	seconds -= minutes * 60
	return fmt.Sprintf("%d hour(s) and %d minute(s)", hours, minutes)
	// return fmt.Sprintf("%d hour(s), %d minute(s) and %d second(s)", hours, minutes, seconds)
}

// AddReminder is a function that adds a reminder for an event to the database
func (db *Database) AddReminder(reminder models.Reminder) error {
	reminders := db.GetCollection(REMINDERS_COLLECTION)
	_, err := reminders.InsertOne(context.TODO(), reminder)
	if err != nil {
		log.Println("error adding reminder: ", err)
		return fmt.Errorf("error adding reminder")
	}

	return nil
}

// Send is a function that sends a remainder for an event to the email address of the user
func (db *Database) SendReminder(r models.Reminder) error {
	from := mail.NewEmail("golangcc", "golangcc42@gmail.com")
	to := mail.NewEmail("", r.Email)
	subject := fmt.Sprintf("Reminder for event: %s", r.Subject)
	plainTextContent := fmt.Sprintf("Hello, your event \"%s\" is about to start in %s.", r.Subject, secondsToString(r.EventStart-r.Time))
	htmlContent := plainTextContent
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	message.SendAt = int(r.Time)

	sendgridApiKey := viper.GetString("SENDGRID_API_KEY")

	client := sendgrid.NewSendClient(sendgridApiKey)
	_, err := client.Send(message)
	if err != nil {
		log.Println("sendgrid error: ", err)
		return err
	}

	// delete reminder after being sent
	err = db.DeleteReminder(r)
	if err != nil {
		log.Println("error with deleting reminder after sending: ", err)
		return err
	}

	log.Println("email sent")
	return nil
}

// SendReminders is a function that goes through all reminders in the database and tries to send them all
func (db *Database) SendReminders() {
	remindersCollection := db.GetCollection(REMINDERS_COLLECTION)
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
				err := db.DeleteReminder(reminder)
				if err != nil {
					log.Println("error deleting reminder: ", err)
				}
			}(reminder)
		} else {
			go func(reminder models.Reminder) {
				err := db.SendReminder(reminder)
				if err != nil {
					log.Println("error sending reminder: ", err)
				}
			}(reminder)
		}
	}
}

// DeleteReminder is a function that deletes a reminder from the database
func (db *Database) DeleteReminder(reminder models.Reminder) error {
	reminders := db.GetCollection(REMINDERS_COLLECTION)

	_, err := reminders.DeleteOne(context.TODO(), reminder)
	if err != nil {
		log.Println("error deleting reminder:", err)
		return err
	}

	log.Println("reminder deleted successfully")
	return nil
}

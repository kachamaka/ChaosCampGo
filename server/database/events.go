package database

import (
	"context"
	"fmt"
	"log"

	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// AddEvent is a function that adds an event to the events array of user in the database
func (db *Database) AddEvent(userID string, event models.Event) error {
	events := db.GetCollection(EVENTS_COLLECTION)
	filter := bson.M{"user_id": userID}
	update := bson.M{"$push": bson.M{"events": event}}

	result := events.FindOneAndUpdate(context.TODO(), filter, update)
	if result.Err() == mongo.ErrNoDocuments {
		userEvents := models.UserEvents{
			UserID: userID,
			Events: []models.Event{event},
		}
		_, err := events.InsertOne(context.TODO(), userEvents)
		if err != nil {
			log.Println("error with adding user events to database: ", err)
			return fmt.Errorf("error adding user events to database")
		}
	} else if result.Err() != nil {
		log.Println("error with adding event: ", result.Err())
		return result.Err()
	}

	return nil
}

// GetEvents is a function that fetches all events for user with ID from the database
func (db *Database) GetEvents(ID string, eventsResponse *models.EventsResponse) error {
	events := db.GetCollection(EVENTS_COLLECTION)
	filter := bson.M{"user_id": ID}

	result := events.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		log.Println("no events for this user")
		return fmt.Errorf("no events for this user")
	} else if result.Err() != nil {
		log.Println("error getting events: ", result.Err())
		return fmt.Errorf("error getting events")
	}

	err := result.Decode(eventsResponse)
	if err != nil {
		log.Println("error decoding events: ", err)
		return fmt.Errorf("error decoding events")
	}

	return nil
}

// DeleteEvent is a function that deletes an event from the events array of user in the database
func (db *Database) DeleteEvent(ID string, event models.Event) error {
	events := db.GetCollection(EVENTS_COLLECTION)
	filter := bson.M{"user_id": ID}
	update := bson.M{"$pull": bson.M{"events": event}}

	result := events.FindOneAndUpdate(context.TODO(), filter, update)
	if result.Err() == mongo.ErrNoDocuments {
		log.Println("no events to delete", result.Err())
		return fmt.Errorf("no events to delete")
	} else if result.Err() != nil {
		log.Println("error with deleting event: ", result.Err())
		return result.Err()
	}

	return nil
}

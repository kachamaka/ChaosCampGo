package database

import (
	"context"
	"fmt"
	"log"

	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func (db *Database) AddEvent(ID string, event models.Event) error {
	events := db.GetCollection(EVENTS_COLLECTION)
	filter := bson.M{"user_id": ID}
	update := bson.M{"$push": bson.M{"events": event}}

	result := events.FindOneAndUpdate(context.TODO(), filter, update)
	if result.Err() == mongo.ErrNoDocuments {
		userEvents := models.UserEvents{
			UserID: ID,
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

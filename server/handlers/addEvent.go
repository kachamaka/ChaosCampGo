package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func AddEventHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	ID, ok := r.Context().Value("_id").(string)
	if !ok {
		encoder.Encode(models.BasicResponse{Success: false, Message: "id from auth header not string"})
		log.Println("err by getting id from auth header")
		return
	}

	var event models.Event
	if err := decoder.Decode(&event); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "can't read body"})
		log.Println("err by decoding body: ", err)
		return
	}

	events := database.Get().GetCollection(database.EVENTS_COLLECTION)
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
			encoder.Encode(models.BasicResponse{Success: false, Message: "error adding user events to database"})
			log.Println("error with adding user events to database: ", err)
			return
		}
	} else if result.Err() != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: result.Err().Error()})
		log.Println("error with adding event: ", result.Err().Error())
		return
	}

	encoder.Encode(models.BasicResponse{Success: true, Message: "event added successfully"})
}

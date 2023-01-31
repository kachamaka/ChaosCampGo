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

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	ID, ok := r.Context().Value("_id").(string)
	if !ok {
		encoder.Encode(models.BasicResponse{Success: false, Message: "id from auth header not string"})
		log.Println("err by getting id from auth header")
		return
	}

	var eventsResponse models.EventsResponse

	events := database.Get().GetCollection(database.EVENTS_COLLECTION)
	filter := bson.M{"_id": ID}
	result := events.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		encoder.Encode(models.BasicResponse{Success: false, Message: "no events for this user"})
		log.Println("no events for this user")
		return
	} else if result.Err() != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "error getting events"})
		log.Println("error getting events: ", result.Err())
		return
	}
	err := result.Decode(&eventsResponse)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "error decoding events"})
		log.Println("error decoding events: ", err)
		return
	}

	eventsResponse.BasicResponse = models.BasicResponse{Success: true, Message: "event added successfully"}
	encoder.Encode(eventsResponse)
}

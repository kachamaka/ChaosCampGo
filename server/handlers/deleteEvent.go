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

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
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
	update := bson.M{"$pull": bson.M{"events": event}}
	result := events.FindOneAndUpdate(context.TODO(), filter, update)
	if result.Err() == mongo.ErrNoDocuments {
		encoder.Encode(models.BasicResponse{Success: true, Message: "no events to delete"})
		log.Println("no events to delete")
		return
	} else if result.Err() != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: result.Err().Error()})
		log.Println("error with deleting event: ", result.Err().Error())
		return
	}

	encoder.Encode(models.BasicResponse{Success: true, Message: "event deleted successfully"})
}

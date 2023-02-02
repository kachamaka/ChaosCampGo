package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func AddReminderHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	headerID, ok := r.Context().Value("_id").(string)
	if !ok {
		encoder.Encode(models.BasicResponse{Success: false, Message: "id from auth header not string"})
		log.Println("err by getting id from auth header")
		return
	}
	ID, err := primitive.ObjectIDFromHex(headerID)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "can't convert id to objectid"})
		log.Println("err converting to objectid: ", err)
		return
	}

	var req models.ReminderRequest
	if err := decoder.Decode(&req); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "can't read body"})
		log.Println("err by decoding body: ", err)
		return
	}

	users := database.Get().GetCollection(database.USERS_COLLECTION)
	filter := bson.M{"_id": ID}
	result := users.FindOne(context.TODO(), filter)
	if result.Err() != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "error getting user"})
		log.Println("err getting user: ", result.Err())
		return
	}

	var user models.User
	err = result.Decode(&user)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "error decoding user"})
		log.Println("err decoding user: ", err)
		return
	}

	reminder := models.Reminder{
		UserID:    headerID,
		Email:     user.Email,
		Subject:   req.Subject,
		Time:      req.StartTime - req.TimeAhead,
		StartTime: req.StartTime,
	}
	reminders := database.Get().GetCollection(database.REMINDERS_COLLECTION)
	_, err = reminders.InsertOne(context.TODO(), reminder)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "error adding reminder"})
		log.Println("err adding reminder: ", err)
		return
	}

	// reminder.Time = time.Now().Add(time.Second * 30).Unix()
	go reminder.Send()

	encoder.Encode(models.BasicResponse{Success: true, Message: "reminder added successfully"})
}

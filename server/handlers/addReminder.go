package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
)

func AddReminderHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	stringID, objectID, err := database.GetHeaders(r)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error()})
		log.Println("err getting headers: ", err)
		return
	}

	var req models.ReminderRequest
	if err := decoder.Decode(&req); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "can't read body"})
		log.Println("err by decoding body: ", err)
		return
	}

	user, err := database.GetUser(objectID)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error()})
		log.Println("err getting user: ", err)
		return
	}

	reminder := models.Reminder{
		UserID:    stringID,
		Email:     user.Email,
		Subject:   req.Subject,
		Time:      req.StartTime - req.TimeAhead,
		StartTime: req.StartTime,
	}

	err = database.Get().AddReminder(reminder)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error()})
		log.Println("err adding reminder: ", err)
		return
	}

	// reminder.Time = time.Now().Add(time.Second * 30).Unix()

	//send reminder
	go database.Send(reminder)

	encoder.Encode(models.BasicResponse{Success: true, Message: "reminder added successfully"})
}

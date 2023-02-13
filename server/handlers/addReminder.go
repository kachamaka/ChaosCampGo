package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
	"github.com/kachamaka/chaosgo/status"
)

// AddReminderHandler is a function that adds a reminder for an event to the database
func AddReminderHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		encoder.Encode(models.BasicResponse{Success: false, Message: "method not allowed", Status: status.METHOD_ERROR})
		log.Println("method not allowed")
		return
	}

	stringID, objectID, err := database.GetHeaders(r)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error(), Status: status.AUTHORIZATION_ERROR})
		log.Println("error by get headers: ", err)
		return
	}

	var req models.ReminderRequest
	if err := decoder.Decode(&req); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "can't read body", Status: status.BODY_ERROR})
		log.Println("error by decode body: ", err)
		return
	}

	user, err := database.Get().GetUserByID(objectID)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error(), Status: status.GET_USER_ERROR})
		log.Println("error by get user: ", err)
		return
	}

	reminder := models.Reminder{
		UserID:     stringID,
		Email:      user.Email,
		Subject:    req.Subject,
		Time:       req.EventStart - req.TimeAhead,
		EventStart: req.EventStart,
	}

	err = database.Get().AddReminder(reminder)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error(), Status: status.ADD_REMINDER_ERROR})
		log.Println("error by add reminder: ", err)
		return
	}

	//send reminder
	// fmt.Println(reminder)
	go func(remider models.Reminder) {
		err := database.Get().SendReminder(reminder)
		if err != nil {
			log.Println("error by send reminder: ", err)
			return
		}
	}(reminder)

	encoder.Encode(models.BasicResponse{Success: true, Message: "reminder added successfully", Status: status.OK})
}

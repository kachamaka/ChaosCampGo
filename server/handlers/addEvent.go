package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
	"github.com/kachamaka/chaosgo/status"
)

// AddEventHandler is a function that adds an event to the database
func AddEventHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		encoder.Encode(models.BasicResponse{Success: false, Message: "method not allowed", Status: status.METHOD_ERROR})
		log.Println("method not allowed")
		return
	}

	stringID, err := database.GetHeader(r)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error(), Status: status.AUTHORIZATION_ERROR})
		log.Println("error by get header", err)
		return
	}

	var event models.Event
	if err := decoder.Decode(&event); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "can't read body", Status: status.BODY_ERROR})
		log.Println("error by decode body: ", err)
		return
	}

	err = database.Get().AddEvent(stringID, event)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error(), Status: status.ADD_EVENT_ERROR})
		log.Println("error by add event: ", err)
		return
	}

	encoder.Encode(models.BasicResponse{Success: true, Message: "event added successfully", Status: status.OK})
}

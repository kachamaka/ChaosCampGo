package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
	"github.com/kachamaka/chaosgo/status"
)

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	stringID, err := database.GetHeader(r)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "id from auth header not string", Status: status.AUTHORIZATION_ERROR})
		log.Println("error by get header")
		return
	}

	var event models.Event
	if err := decoder.Decode(&event); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "can't read body", Status: status.BODY_ERROR})
		log.Println("error by decode body: ", err)
		return
	}

	err = database.Get().DeleteEvent(stringID, event)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error(), Status: status.DELETE_EVENT_ERROR})
		log.Println("error deleting event: ", err)
		return
	}

	encoder.Encode(models.BasicResponse{Success: true, Message: "event deleted successfully", Status: status.OK})
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
	"github.com/kachamaka/chaosgo/status"
)

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	stringID, err := database.GetHeader(r)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error(), Status: status.AUTHORIZATION_ERROR})
		log.Println("error by get header", err)
		return
	}

	var eventsResponse models.EventsResponse
	err = database.Get().GetEvents(stringID, &eventsResponse)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error(), Status: status.GET_EVENTS_ERROR})
		log.Println("error by get events", err)
		return
	}

	eventsResponse.BasicResponse = models.BasicResponse{Success: true, Message: "events fetched successfully", Status: status.OK}
	encoder.Encode(eventsResponse)
}

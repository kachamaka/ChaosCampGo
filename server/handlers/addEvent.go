package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kachamaka/chaosgo/models"
)

type Event struct {
	UserID  string `json:"userID" bson:"userID"`
	Subject string `json:"subject" bson:"subject"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	ID := r.Context().Value("_id")
	fmt.Println(ID)

	_ = decoder

	encoder.Encode(models.BasicResponse{Success: true, Message: "all good"})
}

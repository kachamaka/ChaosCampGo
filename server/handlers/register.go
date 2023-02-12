package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
	"github.com/kachamaka/chaosgo/status"
	"github.com/kachamaka/chaosgo/tokens"
)

// RegisterHandler is a function that register a new user to the database
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		encoder.Encode(models.BasicResponse{Success: false, Message: "method not allowed", Status: status.METHOD_ERROR})
		log.Println("method not allowed")
		return
	}

	var request models.RegisterRequest
	if err := decoder.Decode(&request); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "err by decoding", Status: status.AUTHORIZATION_ERROR})
		log.Println("error by decode body: ", err)
		return
	}

	ID, err := database.Get().Register(request)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error(), Status: status.REGISTER_ERROR})
		log.Println("error by register: ", err)
		return
	}

	token, err := tokens.GenerateToken(ID)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "couldn't generate token", Status: status.TOKEN_ERROR})
		log.Println("error by generate token: ", err)
		return
	}

	encoder.Encode(models.AuthResponse{Token: token, BasicResponse: models.BasicResponse{Success: true, Message: "register successful", Status: status.OK}})
}

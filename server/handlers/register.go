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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// if method POST
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

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

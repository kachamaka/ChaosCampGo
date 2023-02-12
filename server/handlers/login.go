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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	var request models.LoginRequest
	if err := decoder.Decode(&request); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "can't read body", Status: status.BODY_ERROR})
		log.Println("error by decode body: ", err)
		return
	}

	ID, err := database.Get().Login(request)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error(), Status: status.LOGIN_ERROR})
		log.Println("error by login", err)
		return
	}

	token, err := tokens.GenerateToken(ID)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "couldn't generate token", Status: status.TOKEN_ERROR})
		log.Println("error by generate token", err)
		return
	}

	encoder.Encode(models.AuthResponse{Token: token, BasicResponse: models.BasicResponse{Success: true, Message: "login successful", Status: status.OK}})
}

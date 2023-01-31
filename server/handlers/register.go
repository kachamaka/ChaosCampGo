package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
	"github.com/kachamaka/chaosgo/tokens"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// if method POST
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	var req models.RegisterRequest

	if err := decoder.Decode(&req); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "err by decoding"})
		log.Println("err by decoding body: ", err)
		return
	}

	usernameExists, err := req.UsernameExists(req.Username)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "error checking if user exists"})
		log.Println("error checking if user exists: ", err)
		return
	}
	if usernameExists {
		encoder.Encode(models.BasicResponse{Success: false, Message: "user already exists"})
		log.Println("error user exists")
		return
	}

	if err = req.HashPassword(); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "error hashing password"})
		log.Println("error hashing password: ", err)
		return
	}

	users := database.Get().GetCollection(database.USERS_COLLECTION)
	result, err := users.InsertOne(r.Context(), req)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error()})
		log.Println("error with register: ", err)
		return
	}

	id := fmt.Sprintf("%v", result.InsertedID)
	token, err := tokens.GenerateToken(id)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "couldn't generate token"})
		log.Println("error with generating token: ", err)
		return
	}

	encoder.Encode(models.AuthResponse{Token: token, BasicResponse: models.BasicResponse{Success: true, Message: "all good"}})
}

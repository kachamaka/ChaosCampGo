package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/argon2"
	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
)

type RegisterRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type RegisterResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires"`
	models.BasicResponse
}

func (req *RegisterRequest) HashPassword() {
	hashed, err := argon2.GenerateFromPassword(req.Password)
	if err != nil {
		log.Fatal("error hashing password")
		return
	}
	req.Password = hashed
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	var req RegisterRequest
	err := decoder.Decode(&req)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "err by decoding"})
		log.Println("err by decoding body: ", err)
		return
	}

	req.HashPassword()

	col := database.Get().GetCollection("users")

	_, err = col.InsertOne(r.Context(), req)
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error()})
		log.Println("can't save user: ", err)
		return
	}

	encoder.Encode(models.BasicResponse{Success: true, Message: "all good"})
}

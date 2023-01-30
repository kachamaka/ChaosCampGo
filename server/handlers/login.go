package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
	"github.com/kachamaka/chaosgo/tokens"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (loginReq *LoginRequest) MatchPassword() (bson.ObjectId, error) {
	users := database.Get().GetCollection(database.USERS_COLLECTION)

	var user models.User
	filter := bson.M{"username": loginReq.Username}
	result := users.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		return "", fmt.Errorf("no users in collection")
	}
	err := result.Decode(&user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return "", fmt.Errorf("error comparing password and hash")
	}

	return user.ID, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	var req LoginRequest

	if err := decoder.Decode(&req); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "can't read body"})
		log.Println("err by decoding body: ", err)
		return
	}

	ID, err := req.MatchPassword()
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "passwords don't match"})
		log.Println("err by passwords mismatch: ", err)
		return
	}
	token, err := tokens.GenerateToken(ID.Hex())
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "couldn't generate token"})
		log.Println("err with generating the token")
		return
	}

	encoder.Encode(models.AuthResponse{Token: token, BasicResponse: models.BasicResponse{Success: true, Message: "all good"}})
}

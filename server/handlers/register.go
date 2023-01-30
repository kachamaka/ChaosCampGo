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

type RegisterRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

func (registerReq *RegisterRequest) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), 14)
	if err != nil {
		return err
	}
	registerReq.Password = string(bytes)
	return nil
}

func (registerReq *RegisterRequest) UsernameExists(username string) (bool, error) {
	users := database.Get().GetCollection(database.USERS_COLLECTION)

	var user models.User
	filter := bson.M{"username": username}
	result := users.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		return false, nil
	}
	err := result.Decode(&user)
	if err != nil {
		return false, err
	}

	if user != (models.User{}) {
		return true, nil
	}

	return false, nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// if method POST
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	var req RegisterRequest

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
		log.Println("error user exists: ", err)
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

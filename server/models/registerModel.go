package models

import (
	"context"

	"github.com/kachamaka/chaosgo/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

	var user User
	filter := bson.M{"username": username}
	result := users.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		return false, nil
	}
	err := result.Decode(&user)
	if err != nil {
		return false, err
	}

	if user != (User{}) {
		return true, nil
	}

	return false, nil
}

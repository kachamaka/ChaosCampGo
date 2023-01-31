package models

import (
	"context"
	"fmt"

	"github.com/kachamaka/chaosgo/database"
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

	var user User
	filter := bson.M{"username": loginReq.Username}
	result := users.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		return "", fmt.Errorf("no users in collection")
	}
	err := result.Decode(&user)
	if err != nil {
		return "", err
	}
	if user == (User{}) {
		return "", fmt.Errorf("no such user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return "", fmt.Errorf("error comparing password and hash")
	}

	return user.ID, nil
}

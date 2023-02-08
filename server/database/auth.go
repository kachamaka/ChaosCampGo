package database

import (
	"context"
	"fmt"
	"log"

	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func (db *Database) Login(request models.LoginRequest) (string, error) {
	users := db.GetCollection(USERS_COLLECTION)

	var user models.User
	filter := bson.M{"username": request.Username}
	result := users.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		return "", fmt.Errorf("no users in collection")
	}
	err := result.Decode(&user)
	if err != nil {
		return "", err
	}
	if user == (models.User{}) {
		return "", fmt.Errorf("no such user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", fmt.Errorf("error comparing password and hash")
	}

	return user.ID, nil
}

func (db *Database) Register(request models.RegisterRequest) (string, error) {
	usernameExists, err := db.UsernameExists(request.Username)
	if err != nil {
		log.Println("error checking if user exists: ", err)
		return "", fmt.Errorf("error checking if user exists")
	}
	if usernameExists {
		log.Println("error user exists")
		return "", fmt.Errorf("user already exists")
	}

	if err = request.HashPassword(); err != nil {
		log.Println("error hashing password: ", err)
		return "", fmt.Errorf("error hashing password")
	}

	users := db.GetCollection(USERS_COLLECTION)
	result, err := users.InsertOne(context.TODO(), request)
	if err != nil {
		log.Println("error with register: ", err)
		return "", fmt.Errorf("error registering user")
	}

	ID := fmt.Sprintf("%v", result.InsertedID)
	// ID, ok := result.InsertedID.(string)
	// if !ok {
	// 	log.Println("error with type asserting InsertedID: ")
	// 	return "", fmt.Errorf("error getting registered user ID")
	// }

	return ID, nil
}

func (db *Database) UsernameExists(username string) (bool, error) {
	users := db.GetCollection(USERS_COLLECTION)

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

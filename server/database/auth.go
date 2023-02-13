package database

import (
	"context"
	"fmt"
	"log"

	"github.com/kachamaka/chaosgo/models"
	"golang.org/x/crypto/bcrypt"
)

// Login is a function that finds user with username in database and ensures the passwords match, returns token string and error
func (db *Database) Login(request models.LoginRequest) (string, error) {
	user, err := db.GetUser(request.Username)
	if err != nil {
		return "", err
	}
	if user == (models.User{}) {
		return "", fmt.Errorf("no such user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", fmt.Errorf("error comparing password and hash: %v", err)
	}

	return user.ID, nil
}

// Register is a function that checks if username is unique, hashes the password and inserts
// the newly created user instance to the database, returns token string and error
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

	return ID, nil
}

// UsernameExists is a function that checks if a username is already present in the users collection in the database
func (db *Database) UsernameExists(username string) (bool, error) {
	user, err := db.GetUser(username)
	if err != nil {
		return false, err
	}

	if user != (models.User{}) {
		return true, nil
	}

	return false, nil
}

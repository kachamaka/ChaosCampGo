package database

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// GetHeader is a function that gets the '_id' property from a header
func GetHeader(r *http.Request) (string, error) {
	headerID, ok := r.Context().Value("_id").(string)
	if !ok {
		return "", fmt.Errorf("id from auth header not string")
	}
	return headerID, nil
}

// GetHeaders is a function that gets the '_id' property from a header and also converts it to the
// primitive.ObjectID type and returns both instances as they are almost always needed
func GetHeaders(r *http.Request) (string, primitive.ObjectID, error) {
	stringID, err := GetHeader(r)
	if err != nil {
		return "", primitive.NilObjectID, err
	}
	objectID, err := primitive.ObjectIDFromHex(stringID)
	if err != nil {
		return "", primitive.NilObjectID, fmt.Errorf("can't convert id to objectid")
	}
	return stringID, objectID, err
}

// GetUser is a function that gets a user with specified ID from the database
func (db *Database) GetUserByID(ID primitive.ObjectID) (models.User, error) {
	users := db.GetCollection(USERS_COLLECTION)
	filter := bson.M{"_id": ID}
	result := users.FindOne(context.TODO(), filter)
	if result.Err() != nil {
		return models.User{}, result.Err()
	}

	var user models.User
	err := result.Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// GetUser is a function that gets a user with specified username from database
func (db *Database) GetUser(username string) (models.User, error) {
	users := db.GetCollection(USERS_COLLECTION)
	filter := bson.M{"username": username}
	result := users.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		return models.User{}, nil
	}
	if result.Err() != nil {
		return models.User{}, result.Err()
	}

	var user models.User
	err := result.Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

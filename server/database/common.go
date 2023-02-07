package database

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func GetHeader(r *http.Request) (string, error) {
	headerID, ok := r.Context().Value("_id").(string)
	if !ok {
		return "", fmt.Errorf("id from auth header not string")
	}
	return headerID, nil
}

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

func GetUser(ID primitive.ObjectID) (models.User, error) {
	users := Get().GetCollection(USERS_COLLECTION)
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

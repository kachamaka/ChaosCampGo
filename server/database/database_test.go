package database_test

import (
	"context"
	"log"
	"testing"

	"github.com/kachamaka/chaosgo/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = database.Database{
	Config: database.Config{
		ServerAddress:   "",
		DatabaseAddress: "mongodb://localhost:27017",
		DatabaseName:    "mockgocc",
	},
}

func NewMockDatabase() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	session := client.Database("mockgocc")
	return session, nil
}

func TestRegister(t *testing.T) {
	// mockCol := &mockCollection{}
	// res, err :=
}

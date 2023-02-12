package database_test

import (
	"context"
	"log"
	"testing"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
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
	// err = InitializeCollections(session)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }

	return session, nil
}

func InitializeCollections(db *mongo.Database) error {
	err := db.CreateCollection(context.TODO(), "users")
	if err != nil {
		return err
	}
	err = db.CreateCollection(context.TODO(), "events")
	if err != nil {
		return err
	}
	err = db.CreateCollection(context.TODO(), "reminders")
	if err != nil {
		return err
	}

	return nil
}

func TestRegister(t *testing.T) {
	// _, err := NewMockDatabase()
	// if err != nil {
	// 	t.Error("db error", err)
	// }
	db.Connect()

	req := models.RegisterRequest{
		Username: "user1",
		Password: "pass",
	}
	req.HashPassword()
	_, err := db.Register(req)
	if err != nil {
		t.Error("register error", err)
	}

	// mockCol := &mockCollection{}
	// res, err :=
}

package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

type Database struct {
	db *mongo.Database
}

var conn *Database

func Get() *Database {
	if conn == nil {
		// Create the connection
		conn = &Database{
			// Connection details
		}
	}
	return conn
}

func (db *Database) GetCollection(col string) *mongo.Collection {
	return db.db.Collection(col)
}

func (db *Database) Connect() {
	session, err := connect()
	if err != nil {
		log.Fatal("Error connecting to DB")
		return
	}

	db.db = session
}

func (db *Database) Disconnect() {
	db.db.Client().Disconnect(context.TODO())
}

func connect() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//comment?
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	session := client.Database("chaosgo")
	return session, nil
}

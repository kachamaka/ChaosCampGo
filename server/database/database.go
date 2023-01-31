package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI = "mongodb://localhost:27017"

// users collection in database
const USERS_COLLECTION = "users"
const EVENTS_COLLECTION = "events"

const DATABASE = "chaosgo"

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

func (db *Database) GetCollection(collection string) *mongo.Collection {
	return db.db.Collection(collection)
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
	clientOptions := options.Client().ApplyURI(URI)
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

	session := client.Database(DATABASE)
	return session, nil
}

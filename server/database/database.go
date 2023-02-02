package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const HTTP_ADDRESS = "0.0.0.0:8888"
// const URI = "mongodb://localhost:27017"

// collections in database
const USERS_COLLECTION = "users"
const EVENTS_COLLECTION = "events"
const REMINDERS_COLLECTION = "reminders"

const DATABASE = "chaosgo"

type Database struct {
	Config Config
	db     *mongo.Database
}

var conn *Database

func Get() *Database {
	if conn == nil {
		// Create the connection
		config, err := LoadConfig(".")
		if err != nil {
			log.Fatalf("error with config, %v", err)
		}

		conn = &Database{
			Config: *config,
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
	clientOptions := options.Client().ApplyURI(conn.Config.DatabaseAddress)
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

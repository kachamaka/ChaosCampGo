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

func GetInstance() *Database {
	if conn == nil {
		// Create the connection
		conn = &Database{
			// Connection details
		}
	}
	return conn
}

func (d *Database) Connect() {
	d.db = connect()
}

func connect() *mongo.Database {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//comment?
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("chaosgo")
	return db
}

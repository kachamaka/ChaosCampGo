package database

import (
	"context"
	"log"

	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// collections in database
const USERS_COLLECTION = "users"
const EVENTS_COLLECTION = "events"
const REMINDERS_COLLECTION = "reminders"

// DatabaseSchema is an interface model for the functionality that the database supports
type DatabaseSchema interface {
	Connect()
	Disconnect()

	Login(models.LoginRequest) (string, error)
	Register(models.RegisterRequest) (string, error)
	UsernameExists(string) (bool, error)

	AddEvent(string, models.Event) error
	GetEvents(string, *models.EventsResponse) error
	DeleteEvent(string, models.Event) error

	AddReminder(models.Reminder) error
}

// Database is a struct model for the database
type Database struct {
	Config Config
	db     *mongo.Database
}

// conn is a Database instance
var conn *Database

// assertion that Database struct implements DatabaseSchema interface
var _ DatabaseSchema = (*Database)(nil)

// Get is a function that implements the Singleton design pattern and returns the only existing
// instance of the database or creates it if such does not exist
func Get() *Database {
	if conn == nil {
		// Load config properties
		config, err := LoadConfig(".")
		if err != nil {
			log.Fatalf("error with config, %v", err)
		}

		// Create the connection
		conn = &Database{
			Config: *config,
		}
	}
	return conn
}

// GetCollection is a function that returns a collection instance for a collection in the database
func (db *Database) GetCollection(collection string) *mongo.Collection {
	return db.db.Collection(collection)
}

// Connect is a function that establishes a connection with the database
func (db *Database) Connect() {
	session, err := connect(db.Config.DatabaseName)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
		return
	}

	db.db = session
}

// Disconnect is a function that disconnects from the database
func (db *Database) Disconnect() {
	db.db.Client().Disconnect(context.TODO())
}

// connect is a function that connects with the database, attemps to ping it and returns the session
func connect(databaseName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(conn.Config.DatabaseAddress)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("err connecting to db:", err)
		return nil, err
	}

	//comment?
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("err pinging db:", err)
		return nil, err
	}

	session := client.Database(databaseName)
	return session, nil
}

package database

import (
	"context"
	"log"

	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// collections in database
const USERS_COLLECTION = "users"
const EVENTS_COLLECTION = "events"
const REMINDERS_COLLECTION = "reminders"

// DatabaseSchema is an interface model for the functionality that the database supports
type DatabaseSchema interface {
	connect(string, string) (*mongo.Database, error)
	Connect(string, string)
	Disconnect()

	Login(models.LoginRequest) (string, error)
	Register(models.RegisterRequest) (string, error)
	UsernameExists(string) (bool, error)
	GetUserByID(ID primitive.ObjectID) (models.User, error)
	GetUser(username string) (models.User, error)

	AddEvent(string, models.Event) error
	GetEvents(string, *models.EventsResponse) error
	DeleteEvent(string, models.Event) error

	AddReminder(models.Reminder) error
	DeleteReminder(reminder models.Reminder) error
	SendReminder(r models.Reminder) error
	SendReminders()
}

// Database is a struct model for the database
type Database struct {
	db *mongo.Database
}

// conn is a Database instance
var conn *Database

// assertion that Database struct implements DatabaseSchema interface
var _ DatabaseSchema = (*Database)(nil)

// Get is a function that implements the Singleton design pattern and returns the only existing
// instance of the database or creates it if such does not exist
func Get() *Database {
	if conn == nil {
		// Create the connection
		conn = &Database{}
	}
	return conn
}

// GetCollection is a function that returns a collection instance for a collection in the database
func (db *Database) GetCollection(collection string) *mongo.Collection {
	return db.db.Collection(collection)
}

// Connect is a function that establishes a connection with the database
func (db *Database) Connect(databaseAddress string, databaseName string) {
	session, err := db.connect(databaseAddress, databaseName)
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
func (db *Database) connect(databaseAddress string, databaseName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(databaseAddress)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("err connecting to db:", err)
		return nil, err
	}
	if client == nil {
		return nil, nil
	}

	session := client.Database(databaseName)
	return session, nil
}

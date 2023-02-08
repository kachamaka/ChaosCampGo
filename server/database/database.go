package database

import (
	"context"
	"fmt"
	"log"

	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MAILGUN_DOMAIN = "sandbox0cb6f5028356498690e13f4374a8072f.mailgun.org"
const MAILGUN_PRIVATE_API_KEY = "25a33e6191a65ea68c448539651414d1-d1a07e51-7dd42058"

// collections in database
const USERS_COLLECTION = "users"
const EVENTS_COLLECTION = "events"
const REMINDERS_COLLECTION = "reminders"

const DATABASE = "chaosgo"

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

type Database struct {
	Config Config
	db     *mongo.Database
}

func (db *Database) AddReminder(reminder models.Reminder) error {
	reminders := db.GetCollection(REMINDERS_COLLECTION)
	_, err := reminders.InsertOne(context.TODO(), reminder)
	if err != nil {
		log.Println("err adding reminder: ", err)
		//CUSTOM ERRORS
		return fmt.Errorf("error adding reminder")
	}

	return nil
}

var conn *Database

var _ DatabaseSchema = (*Database)(nil)

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

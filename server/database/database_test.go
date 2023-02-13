package database_test

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/kachamaka/chaosgo/config"
	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func NewMockDatabase() *database.Database {
	if mockDB == nil {
		mockDB = &database.Database{}
		cfg, err := config.LoadConfig("../test")
		if err != nil {
			log.Fatal("test confing file not found")
		}
		testConfig = cfg
	}

	mockDB.Connect(testConfig.DatabaseAddress, testConfig.DatabaseName)
	return mockDB
}

var (
	mockDB     *database.Database
	testConfig *config.Config
	TEST_ID    string = "63ea64ef2e1c7d21f929d50d"

	TEST_USER = models.User{
		ID:       "",
		Username: "user1",
		Password: "pass",
		Email:    "test@test.test",
	}

	TEST_EVENT1 = models.Event{
		Subject: "Test subject1",
		Day:     1,
		Start:   "14:00",
		End:     "17:00",
	}
	TEST_EVENT2 = models.Event{
		Subject: "Test subject2",
		Day:     3,
		Start:   "10:00",
		End:     "12:00",
	}
	TEST_EVENT3 = models.Event{
		Subject: "Test subject3",
		Day:     6,
		Start:   "8:00",
		End:     "17:00",
	}

	TEST_REMINDER1 = models.Reminder{
		UserID:     "",
		Email:      "",
		Subject:    TEST_EVENT1.Subject,
		Time:       time.Now().Add(time.Second).Unix(),
		EventStart: 0,
	}
	TEST_REMINDER2 = models.Reminder{
		UserID:     "",
		Email:      "",
		Subject:    TEST_EVENT2.Subject,
		Time:       time.Now().Add(time.Minute).Unix(),
		EventStart: 0,
	}
	TEST_REMINDER3 = models.Reminder{
		UserID:     "",
		Email:      "",
		Subject:    TEST_EVENT3.Subject,
		Time:       time.Now().Add(-time.Minute).Unix(),
		EventStart: 0,
	}
)

func TestGetHeader(t *testing.T) {
	req, _ := http.NewRequest("", "", nil)
	ctx := context.WithValue(req.Context(), "_id", TEST_ID)
	req = req.WithContext(ctx)

	_, err := database.GetHeader(req)
	if err != nil {
		t.Error("get header error", err)
	}
}

func TestGetHeaders(t *testing.T) {
	req, _ := http.NewRequest("", "", nil)
	ctx := context.WithValue(req.Context(), "_id", TEST_ID)
	req = req.WithContext(ctx)

	_, _, err := database.GetHeaders(req)
	if err != nil {
		t.Error("get headers error", err)
	}
}

func TestRegister(t *testing.T) {
	db := NewMockDatabase()

	req := models.RegisterRequest{
		Username: TEST_USER.Username,
		Password: TEST_USER.Password,
		Email:    TEST_USER.Email,
	}

	_, err := db.Register(req)
	if err != nil {
		t.Error("register error", err)
	}

	_, err = db.Register(req)
	if err == nil {
		t.Error("register same user no error", err)
	}
}

func TestLogin(t *testing.T) {
	db := NewMockDatabase()

	req := models.LoginRequest{
		Username: TEST_USER.Username,
		Password: TEST_USER.Password,
	}
	_, err := db.Login(req)
	if err != nil {
		t.Error("login error", err)
	}

	req.Password = ""
	_, err = db.Login(req)
	if err == nil {
		t.Error("shouldn't have found user wrong pass")
	}

	req.Username = ""
	_, err = db.Login(req)
	if err == nil {
		t.Error("shouldn't have found user wrong username")
	}
}

func TestGetUser(t *testing.T) {
	db := NewMockDatabase()

	_, err := db.GetUser(TEST_USER.Username)
	if err != nil {
		t.Error("get user error", err)
	}
}

func TestGetUserByID(t *testing.T) {
	db := NewMockDatabase()

	user, err := db.GetUser(TEST_USER.Username)
	if err != nil {
		t.Error("get user error", err)
	}

	ID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		t.Error("error casting hex to objectID")
	}
	userByID, err := db.GetUserByID(ID)
	if err != nil {
		t.Error("get user by id error", err)
	}
	if user.ID != userByID.ID || user.Username != userByID.Username {
		t.Error("users not the same")
	}
}

func TestUsernameExists(t *testing.T) {
	db := NewMockDatabase()

	_, err := db.UsernameExists(TEST_USER.Username)
	if err != nil {
		t.Error("username exists error", err)
	}
}

func TestAddEvent(t *testing.T) {
	db := NewMockDatabase()

	user, err := db.GetUser(TEST_USER.Username)
	if err != nil {
		t.Error("get user error", err)
	}

	err = mockDB.AddEvent(user.ID, TEST_EVENT1)
	if err != nil {
		t.Error("add event error", err)
	}
	err = mockDB.AddEvent(user.ID, TEST_EVENT2)
	if err != nil {
		t.Error("add event error", err)
	}
	err = mockDB.AddEvent(user.ID, TEST_EVENT3)
	if err != nil {
		t.Error("add event error", err)
	}
}

func TestDeleteEvent(t *testing.T) {
	db := NewMockDatabase()

	user, err := db.GetUser(TEST_USER.Username)
	if err != nil {
		t.Error("get user error", err)
	}

	err = db.DeleteEvent(user.ID, TEST_EVENT1)
	if err != nil {
		t.Error("delete event error", err)
	}
	err = db.DeleteEvent(user.ID, TEST_EVENT2)
	if err != nil {
		t.Error("delete event error", err)
	}
	err = db.DeleteEvent(user.ID, TEST_EVENT3)
	if err != nil {
		t.Error("delete event error", err)
	}

}

func TestGetEvents(t *testing.T) {
	db := NewMockDatabase()

	user, err := db.GetUser(TEST_USER.Username)
	if err != nil {
		t.Error("get user error", err)
	}

	var eventsResponse models.EventsResponse
	err = db.GetEvents(user.ID, &eventsResponse)
	if err != nil {
		t.Error("get events error", err)
	}

	db.AddEvent(user.ID, TEST_EVENT1)
	db.AddEvent(user.ID, TEST_EVENT2)
	db.AddEvent(user.ID, TEST_EVENT3)

	err = db.GetEvents(user.ID, &eventsResponse)
	if err != nil {
		t.Error("get events error", err)
	}

	if len(eventsResponse.Events) != 3 {
		t.Error("error with events number")
	}
}

func TestAddReminder(t *testing.T) {
	db := NewMockDatabase()

	user, err := db.GetUser(TEST_USER.Username)
	if err != nil {
		t.Error("get user error", err)
	}

	TEST_REMINDER1.UserID = user.ID
	TEST_REMINDER1.Email = user.Email
	err = db.AddReminder(TEST_REMINDER1)
	if err != nil {
		t.Error("add reminder error", err)
	}

	TEST_REMINDER2.UserID = user.ID
	TEST_REMINDER2.Email = user.Email
	err = db.AddReminder(TEST_REMINDER2)
	if err != nil {
		t.Error("add reminder error", err)
	}

	TEST_REMINDER3.UserID = user.ID
	TEST_REMINDER3.Email = user.Email
	err = db.AddReminder(TEST_REMINDER3)
	if err != nil {
		t.Error("add reminder error", err)
	}

}

func TestSendReminder(t *testing.T) {
	db := NewMockDatabase()

	user, err := db.GetUser(TEST_USER.Username)
	if err != nil {
		t.Error("get user error", err)
	}

	TEST_REMINDER1.UserID = user.ID
	TEST_REMINDER1.Email = user.Email
	err = db.SendReminder(TEST_REMINDER1)
	if err != nil {
		t.Error("send reminder error", err)
	}

}

func TestSendReminders(t *testing.T) {
	db := NewMockDatabase()

	db.SendReminders()
}

func TestEmptyCollections(t *testing.T) {
	db := NewMockDatabase()

	col := db.GetCollection(database.USERS_COLLECTION)
	_, err := col.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		t.Error("1.empty collection error", err)
	}
	col = db.GetCollection(database.EVENTS_COLLECTION)
	_, err = col.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		t.Error("2.empty collection error", err)
	}
	col = db.GetCollection(database.REMINDERS_COLLECTION)
	_, err = col.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		t.Error("3.empty collection error", err)
	}
}

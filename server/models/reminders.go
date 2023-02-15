package models

// Reminder is a struct model for reminders of events in the database
type Reminder struct {
	UserID     string `json:"user_id" bson:"user_id"`
	Email      string `json:"email" bson:"email"`
	Subject    string `json:"subject" bson:"subject"`
	Time       int64  `json:"time" bson:"time"`
	EventStart int64  `json:"event_start" bson:"event_start"`
}

// ReminderRequest is a struct model for the JSON body request in the addReminder handler
type ReminderRequest struct {
	Subject    string `json:"subject"`
	EventStart int64  `json:"eventStart"`
	TimeAhead  int64  `json:"timeAhead"` //how many minutes earlier to send a reminder
}

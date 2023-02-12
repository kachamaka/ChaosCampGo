package models

type Reminder struct {
	UserID     string `json:"user_id" bson:"user_id"`
	Email      string `json:"email" bson:"email"`
	Subject    string `json:"subject" bson:"subject"`
	Time       int64  `json:"time" bson:"time"`
	EventStart int64  `json:"eventStart" bson:"eventStart"`
}

type ReminderRequest struct {
	Subject    string `json:"subject"`
	EventStart int64  `json:"eventStart"`
	TimeAhead  int64  `json:"timeAhead"` //how many minutes earlier to send a reminder
}

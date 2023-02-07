package models

type Reminder struct {
	UserID    string `json:"user_id" bson:"user_id"`
	Email     string `json:"email" bson:"email"`
	Subject   string `json:"subject" bson:"subject"`
	Time      int64  `json:"time" bson:"time"`
	StartTime int64  `json:"startTime" bson:"startTime"`
}

type ReminderRequest struct {
	Subject   string `json:"subject"`
	TimeAhead int64  `json:"timeAhead"` //how many minutes earlier to send a reminder
	StartTime int64  `json:"startTime"`
}

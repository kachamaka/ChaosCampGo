package models

const (
	Monday = iota
	Tuesday
	Wednesday
	Thurdsay
	Friday
	Saturday
	Sunday
)

// Event is a struct model for single event
type Event struct {
	Subject string `json:"subject" bson:"subject"`
	Day     int    `json:"day" bson:"day"`
	Start   string `json:"start" bson:"start"`
	End     string `json:"end" bson:"end"`
}

// UserEvents is a struct model for the events in the database
type UserEvents struct {
	UserID string  `bson:"user_id"`
	Events []Event `bson:"events"`
}

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

type Event struct {
	Subject   string `json:"subject" bson:"subject"`
	Day       int    `json:"day" bson:"day"`
	StartTime int64  `json:"startTime" bson:"startTime"`
	Start     string `json:"start" bson:"start"`
	End       string `json:"end" bson:"end"`
}

type UserEvents struct {
	UserID string  `bson:"user_id"`
	Events []Event `bson:"events"`
}

type EventsResponse struct {
	Events []Event `json:"events" bson:"events"`
	BasicResponse
}

package models

type Event struct {
	// UserID  string `json:"userID" bson:"userID"`
	Subject string `json:"subject" bson:"subject"`
	Start   int64  `json:"start" bson:"start"`
	End     int64  `json:"end" bson:"end"`
}

type UserEvents struct {
	ID     string  `bson:"_id"`
	Events []Event `bson:"events"`
}

type EventsResponse struct {
	Events []Event `json:"events" bson:"events"`
	BasicResponse
}

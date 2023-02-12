package models

// BasicResponse is a struct model for the response that is being sent to the frontend
type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// AuthResponse is a struct model for the response that is being sent from the login and register handlers
type AuthResponse struct {
	Token string `json:"token"`
	BasicResponse
}

// EventsResponse is a struct model for the response that is being sent from the getEvents handler
type EventsResponse struct {
	Events []Event `json:"events" bson:"events"`
	BasicResponse
}

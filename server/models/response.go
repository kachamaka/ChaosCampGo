package models

type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AuthResponse struct {
	Token string `json:"token"`
	BasicResponse
}

type EventsResponse struct {
	Events []Event `json:"events" bson:"events"`
	BasicResponse
}

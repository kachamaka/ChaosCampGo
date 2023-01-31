package models

type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type AuthResponse struct {
	Token string `json:"token"`
	BasicResponse
}

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

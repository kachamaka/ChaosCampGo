package models

import "golang.org/x/crypto/bcrypt"

// User is a struct model for the users in the database
type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

// LoginRequest is a struct model for JSON requests for the login handler
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest is a struct model for JSON requests for the register handler
type RegisterRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

// HashPassword is a function for the RegisterRequest struct model which
// hashes the password of the instance
func (registerReq *RegisterRequest) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), 14)
	if err != nil {
		return err
	}
	registerReq.Password = string(bytes)
	return nil
}

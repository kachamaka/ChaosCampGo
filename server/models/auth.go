package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

func (registerReq *RegisterRequest) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), 14)
	if err != nil {
		return err
	}
	registerReq.Password = string(bytes)
	return nil
}

package models

import "gopkg.in/mgo.v2/bson"

type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
}

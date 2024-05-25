package userModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email      string             `json:"email,omitempty"`
	Password   string             `json:"password,omitempty"`
	ProfileUrl string             `json:"profileurl,omitempty"`
	Username   string             `json:"username,omitempty"`
	FullName   string             `json:"fullname,omitempty"`
}

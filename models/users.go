package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model where the profile info is saved
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty", json:"id"`
	Name         string             `bson:"name", json:"name,omitempty"`
	LastName     string             `bson:"lastname", json:"lastname,omitempty"`
	UserBirthday string             `bson:"user_birthday", json:"userBirthday,omitempty"`
	Email        string             `bson:"email", json: "email"`
	Password     string             `bson:"password", json: "password,omitempty"`
	Avatar       string             `bson:"avatar", json: "avatar,omitempty"`
	Biography    string             `bson:"biography", json: "biography,omitempty"`
	WebSite      string             `bson:"webSite", json: "webSite,omitempty"`
}

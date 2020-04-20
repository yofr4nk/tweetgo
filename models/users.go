package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model where the profile info is saved
type User struct {
	ID			primitive.ObjectID	`bson: "_id, omitempty", json: id"`
	Name 		string				`bson: "name", json: name, omitempty"`
	Email 		string				`bson: "email", json: email"`
}

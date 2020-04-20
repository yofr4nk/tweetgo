package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model where the profile info is saved
type User struct {
	ID			primitive.ObjectID	`bson: "_id, omitempty", json: id"`
	Name 		string				`bson: "name", json: name, omitempty"`
	CreatedAt 	time.Time 			`bson: "created_at", json: created_at, omitempty"`
}

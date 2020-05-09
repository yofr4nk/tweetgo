package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Tweet is the model where a news from user is saved
type Tweet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID  string             `bson:"userid" json:"userid,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    int64              `bson:"date" json:"date,omitempty"`
}

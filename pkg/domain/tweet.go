package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tweet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID  string             `bson:"user_id" json:"user_id,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    int64              `bson:"date" json:"date,omitempty"`
}

type Tweets []Tweet

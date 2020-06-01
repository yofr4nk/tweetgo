package saving

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model where the profile info is saved
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name,omitempty"`
	LastName     string             `bson:"lastname" json:"lastname,omitempty"`
	UserBirthday string             `bson:"userbirthday" json:"userbirthday,omitempty"`
	Email        string             `bson:"email" json:"email"`
	Password     string             `bson:"password" json:"password,omitempty"`
	Avatar       string             `bson:"avatar" json:"avatar,omitempty"`
	Banner       string             `bson:"banner" json:"banner,omitempty"`
	Biography    string             `bson:"biography" json:"biography,omitempty"`
	Location     string             `bson:"location" json:"location,omitempty"`
	WebSite      string             `bson:"website" json:"website,omitempty"`
}

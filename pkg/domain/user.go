package domain

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserCtxKey string = "userCtxKey"

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

func SetUserToContext(ctx context.Context, u User) context.Context {
	return context.WithValue(ctx, UserCtxKey, u)
}

func GetUserFromCtx(ctx context.Context) (User, error) {
	if ctx == nil {
		return User{}, errors.New("context is empty")
	}

	r := ctx.Value(UserCtxKey)
	if r != nil {
		return r.(User), nil
	}
	return User{}, errors.New("user not found in context")
}

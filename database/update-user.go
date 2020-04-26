package database

import (
	"context"
	"time"

	"github.com/yofr4nk/tweetgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateUser modify user info
func UpdateUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := DBConnect.Database("tweetgo")
	userCollection := database.Collection("users")

	userToUpdate := make(map[string]interface{})

	if len(u.Name) > 0 {
		userToUpdate["Name"] = u.Name
	}

	if len(u.LastName) > 0 {
		userToUpdate["LastName"] = u.LastName
	}

	userToUpdate["UserBirthday"] = u.UserBirthday

	if len(u.Email) > 0 {
		userToUpdate["Email"] = u.Email
	}

	if len(u.Avatar) > 0 {
		userToUpdate["Avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		userToUpdate["Banner"] = u.Banner
	}

	if len(u.Biography) > 0 {
		userToUpdate["Biography"] = u.Biography
	}

	if len(u.Location) > 0 {
		userToUpdate["Location"] = u.Location
	}

	if len(u.WebSite) > 0 {
		userToUpdate["WebSite"] = u.WebSite
	}

	setContainer := bson.M{
		"$set": userToUpdate,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := userCollection.UpdateOne(ctx, filter, setContainer)

	if err != nil {
		return false, err
	}

	return true, nil

}

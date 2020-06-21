package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/hashing"
)

type UserStorage struct {
	collection *mongo.Collection
}

const (
	userCollection string = "users"
)

func NewUserStorage(db *mongo.Client) *UserStorage {
	database := db.Database(dbName)
	uc := database.Collection(userCollection)

	return &UserStorage{collection: uc}
}

func (storage *UserStorage) SaveUser(u domain.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	u.Password, _ = hashing.HashPassword(u.Password)

	_, err := storage.collection.InsertOne(ctx, u)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (storage *UserStorage) UpdateUser(usrData domain.UserDataContainer, ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	setContainer := bson.M{
		"$set": usrData,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := storage.collection.UpdateOne(ctx, filter, setContainer)
	if err != nil {
		return err
	}

	return nil
}

func (storage *UserStorage) FindUserExists(email string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"email": email}

	count, err := storage.collection.CountDocuments(ctx, condition)
	if err != nil {
		return 0, err
	}

	if count > 0 {
		return count, nil
	}

	return 0, nil
}

func (storage *UserStorage) FindUser(email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"email": email}
	var u domain.User

	err := storage.collection.FindOne(ctx, condition).Decode(&u)
	if err != nil {
		return u, err
	}

	return u, nil
}

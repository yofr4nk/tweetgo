package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/encrypting"
)

type UserStorage struct {
	db *mongo.Client
}

const (
	dbName     string = "tweetgo"
	collection string = "users"
)

func NewUserStorage(db *mongo.Client) *UserStorage {
	return &UserStorage{db: db}
}

func (storage *UserStorage) SaveUser(u domain.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := storage.db.Database(dbName)
	userCollection := database.Collection(collection)

	u.Password, _ = encrypting.HashPassword(u.Password)

	result, err := userCollection.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}

func (storage *UserStorage) FindUserExists(email string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	database := storage.db.Database(dbName)
	userCollection := database.Collection(collection)

	condition := bson.M{"email": email}

	count, err := userCollection.CountDocuments(ctx, condition)

	if err != nil {
		return 0, err
	}

	if count > 0 {
		return count, nil
	}

	return 0, nil
}

package data

import (
	"context"
	"email_auth/initer"
	"email_auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

func InitUserCollection() {
	UserCollection = initer.MongoClient.Database("demo").Collection("users")

	// Set email to the unique index
	UserCollection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.M{"Email": 1},
		Options: options.Index().SetUnique(true),
	})
}

func AddUser(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error) {
	return UserCollection.InsertOne(ctx, user)
}

func GetUserById(ctx context.Context, id any) (*mongo.SingleResult, error) {
	return UserCollection.FindOne(ctx, bson.M{"_id": id}), nil
}

func GetUserByEmail(ctx context.Context, email string) (*mongo.SingleResult, error) {
	return UserCollection.FindOne(ctx, bson.M{"Email": email}), nil
}

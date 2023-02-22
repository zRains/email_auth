package initer

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client
var RedisClient *redis.Client

func initMongo(ctx context.Context) error {
	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(AppConfig.DbUrI)); err != nil {
		return err
	} else {
		MongoClient = client
	}

	if err := MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	fmt.Println("Connected to Mongodb.")

	return nil
}

func initRedis(ctx context.Context) error {
	RedisClient = redis.NewClient(&redis.Options{Addr: AppConfig.RedisUrI, Password: AppConfig.RedisPass, DB: 0})

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		return err
	}

	fmt.Println("Connected to Redis.")

	return nil
}

func InitDatabase(ctx context.Context) {

	if err := initMongo(ctx); err != nil {
		panic(err)
	}

	if err := initRedis(ctx); err != nil {
		panic(err)
	}
}

func CloseDatabase(ctx context.Context) {
	_ = MongoClient.Disconnect(ctx)
	_ = RedisClient.Close()
}

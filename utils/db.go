package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	LoadDotEnv()
}

func CollectionConnection(collection string) *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}
	return client.Database(os.Getenv("DATABASE_NAME")).Collection(collection)
}

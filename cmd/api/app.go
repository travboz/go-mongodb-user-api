package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	Client *mongo.Client
}

func NewApplication() (*application, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, fmt.Errorf("MONGODB_URI is not set")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// verifying connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &application{
		Client: client,
	}, nil
}

func (a *application) Shutdown(ctx context.Context) error {
	return a.Client.Disconnect(ctx)
}

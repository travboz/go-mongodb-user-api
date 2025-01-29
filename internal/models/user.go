package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name,omitempty"`
	Email           string             `bson:"email,omitempty"`
	FavouriteNumber int                `bson:"favourite_number,omitempty"`
	Active          bool               `bson:"active,omitempty"`
}

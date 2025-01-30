package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	Email           string             `json:"email" bson:"email"`
	FavouriteNumber int                `json:"favourite_number" bson:"favourite_number"`
	Active          bool               `json:"active" bson:"active"`
}

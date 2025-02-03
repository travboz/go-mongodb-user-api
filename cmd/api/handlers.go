package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/travboz/backend-projects/go-and-mongo-mohd/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HelloHandler - a helath check handler
func (a *application) HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there, welcome to the Mongo CRUD!")
}

// CreateUser - create a new user
func (a *application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = primitive.NewObjectID()

	collection := a.Client.Database("mongo_user_crud").Collection("users")
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// GetAllUsers - get all users within the collection
func (a *application) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	collection := a.Client.Database("mongo_user_crud").Collection("users")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// GetUserById - fetch a user by a given id
func (a *application) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	req_id := r.PathValue("id")
	id, err := primitive.ObjectIDFromHex(req_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := a.Client.Database("mongo_user_crud").Collection("users")
	var user models.User
	result := collection.FindOne(context.Background(), bson.M{"_id": id})

	if err := result.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// UpdateUser - update a user in the collection
func (a *application) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	req_id := r.PathValue("id")
	id, err := primitive.ObjectIDFromHex(req_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := a.Client.Database("mongo_user_crud").Collection("users")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedUser}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// DeleteUserById - remove a user from a collection by a specific id
func (a *application) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	req_id := r.PathValue("id")

	id, err := primitive.ObjectIDFromHex(req_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := a.Client.Database("mongo_user_crud").Collection("users")
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

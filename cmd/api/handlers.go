package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/travboz/backend-projects/go-and-mongo-mohd/internal/models"
)

// HelloHandler - a helath check handler
func (app *application) HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there, welcome to the Mongo CRUD!")
}

// CreateUser - create a new user
func (app *application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.Storage.Insert(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "user inserted successfully",
	})
}

// GetAllUsers - get all users within the collection
func (app *application) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := app.Storage.FetchAllUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// // GetUserById - fetch a user by a given id
// func (app *application) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
// 	req_id := r.PathValue("id")
// 	id, err := primitive.ObjectIDFromHex(req_id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	collection := app.Client.Database("mongo_user_crud").Collection("users")
// 	var user models.User
// 	result := collection.FindOne(context.Background(), bson.M{"_id": id})

// 	if err := result.Decode(&user); err != nil {
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(user)
// }

// // UpdateUser - update a user in the collection
// func (app *application) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
// 	req_id := r.PathValue("id")
// 	id, err := primitive.ObjectIDFromHex(req_id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	var updatedUser models.User
// 	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	collection := app.Client.Database("mongo_user_crud").Collection("users")
// 	filter := bson.M{"_id": id}
// 	update := bson.M{"$set": updatedUser}

// 	result, err := collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(result)
// }

// // DeleteUserById - remove a user from a collection by a specific id
// func (app *application) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
// 	req_id := r.PathValue("id")

// 	id, err := primitive.ObjectIDFromHex(req_id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	collection := app.Client.Database("mongo_user_crud").Collection("users")
// 	filter := bson.M{"_id": id}
// 	result, err := collection.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(result)
// }

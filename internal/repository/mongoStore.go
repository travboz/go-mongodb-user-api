package repository

import (
	"context"

	"github.com/travboz/backend-projects/go-and-mongo-mohd/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	Client *mongo.Client
}

func NewMongoStore(c *mongo.Client) MongoStore {
	return MongoStore{
		Client: c,
	}
}

func (m *MongoStore) Insert(ctx context.Context, u models.User) error {
	u.ID = primitive.NewObjectID()

	collection := m.Client.Database("mongo_user_crud").Collection("users")
	_, err := collection.InsertOne(ctx, u)

	return err
}

func (m *MongoStore) GetById(ctx context.Context, id string) (*models.User, error) {

	user_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := m.Client.Database("mongo_user_crud").Collection("users")
	var user models.User
	result := collection.FindOne(ctx, bson.M{"_id": user_id})

	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *MongoStore) FetchAllUsers(ctx context.Context) ([]*models.User, error) {
	collection := m.Client.Database("mongo_user_crud").Collection("users")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var users []*models.User
	for cursor.Next(ctx) {
		var user *models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (m *MongoStore) UpdateUser(ctx context.Context, id string, u models.User) error {
	update_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := m.Client.Database("mongo_user_crud").Collection("users")
	filter := bson.M{"_id": update_id}
	update := bson.M{"$set": u}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoStore) DeleteUserById(ctx context.Context, id string) error {
	delete_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := m.Client.Database("mongo_user_crud").Collection("users")
	filter := bson.M{"_id": delete_id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoStore) Shutdown(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

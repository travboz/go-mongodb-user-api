package repository

import (
	"context"

	"github.com/travboz/backend-projects/go-and-mongo-mohd/internal/models"
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
	_, err := collection.InsertOne(context.Background(), u)

	return err
}
func (m *MongoStore) GetById(ctx context.Context, id int64) (*models.User, error) {
	return nil, nil
}
func (m *MongoStore) FetchAllUsers(ctx context.Context) ([]*models.User, error) {
	return nil, nil
}
func (m *MongoStore) UpdateUser(ctx context.Context, id int64, u models.User) error {
	return nil
}
func (m *MongoStore) DeleteUserById(ctx context.Context, id int64) error {
	return nil
}

func (m *MongoStore) Shutdown(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

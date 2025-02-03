package repository

import (
	"context"

	"github.com/travboz/backend-projects/go-and-mongo-mohd/internal/models"
)

type Storage interface {
	Insert(context.Context, models.User) error
	GetById(context.Context, int64) (*models.User, error)
	FetchAllUsers(context.Context) ([]*models.User, error)
	UpdateUser(context.Context, int64, models.User) error
	DeleteUserById(context.Context, int64) error
	Shutdown(context.Context) error
}

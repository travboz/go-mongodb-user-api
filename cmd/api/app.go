package main

import (
	"net/http"

	"github.com/travboz/backend-projects/go-and-mongo-mohd/internal/repository"
)

type application struct {
	Storage repository.Storage
}

func NewApplication(s repository.Storage) *application {
	return &application{
		Storage: s,
	}
}

func (app *application) MountRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.HelloHandler)
	mux.HandleFunc("POST /users", app.CreateUserHandler)
	mux.HandleFunc("GET /users", app.GetAllUsersHandler)
	mux.HandleFunc("GET /users/{id}", app.GetUserByIdHandler)
	// mux.HandleFunc("PUT /users/{id}", app.UpdateUserHandler)
	// mux.HandleFunc("DELETE /users/{id}", app.DeleteUserHandler)

	return mux
}

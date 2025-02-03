package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app, err := NewApplication()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Ensure MongoDB disconnects when the program exits
	defer func() {
		if err := app.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down MongoDB client: %v", err)
		}
	}()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.HelloHandler)
	mux.HandleFunc("POST /users", app.CreateUserHandler)
	mux.HandleFunc("GET /users", app.GetAllUsersHandler)
	mux.HandleFunc("GET /users/{id}", app.GetUserByIdHandler)
	mux.HandleFunc("PUT /users/{id}", app.UpdateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", app.DeleteUserHandler)

	addr := os.Getenv("SERVER_PORT")

	fmt.Printf("Server running at http://localhost%s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/travboz/backend-projects/go-and-mongo-mohd/internal/handlers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HelloHandler)
	mux.HandleFunc("POST /users", handlers.CreateUser)
	mux.HandleFunc("GET /users", handlers.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", handlers.GetUserById)
	mux.HandleFunc("PUT /users/{id}", handlers.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", handlers.DeleteUser)

	addr := os.Getenv("SERVER_PORT")

	fmt.Printf("Server running at http://localhost%s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

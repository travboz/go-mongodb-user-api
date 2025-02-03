package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/travboz/backend-projects/go-and-mongo-mohd/internal/db"
	"github.com/travboz/backend-projects/go-and-mongo-mohd/internal/repository"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}

func main() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}

	mongo, err := db.NewMongoDBClient(uri)
	if err != nil {
		log.Fatal("Unable to instantiate new Mongo client:", err)
	}

	defer mongo.Disconnect(context.Background())

	mongoStore := repository.NewMongoStore(mongo)

	app := NewApplication(&mongoStore)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Ensure MongoDB disconnects when the program exits
	defer func() {
		if err := app.Storage.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down MongoDB client: %v", err)
		}
	}()

	mux := app.MountRoutes()

	srv := &http.Server{
		Addr:         os.Getenv("SERVER_PORT"),
		Handler:      mux,
		WriteTimeout: time.Second * 30, // if our server takes more than 30 seconds to write a response to a client - it times out
		ReadTimeout:  time.Second * 10, // diddo but to read
		IdleTimeout:  time.Minute,      // diddo but when idling
	}

	fmt.Printf("Server running at http://localhost%s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"grpc-api/pkg/auth"
	"grpc-api/pkg/middlewares"
	"grpc-api/pkg/todo"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var err error

	if err = godotenv.Load("../../configs/.env"); err != nil {
		return err
	}

	r := mux.NewRouter()

	authService := auth.SetupRoutes(r, os.Getenv("AUTH_URL"))

	authMiddleware := middlewares.NewAuthMiddleware(authService)

	todo.SetupRoutes(r, os.Getenv("TODO_URL"), authMiddleware)

	if err = http.ListenAndServe(":8080", r); err != nil {
		return err
	}

	return nil
}

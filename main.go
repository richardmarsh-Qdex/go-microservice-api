package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go-microservice-api/internal/handlers"
	"go-microservice-api/internal/middleware"
	"go-microservice-api/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.Connect()

	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/api/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/api/auth/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/auth/login", handlers.Login).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	api.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	api.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	api.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	api.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
	api.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	api.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	api.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	api.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	api.HandleFunc("/orders/{id}", handlers.GetOrder).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

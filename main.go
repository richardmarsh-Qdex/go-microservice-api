package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go-microservice-api/internal/auth"
	"go-microservice-api/internal/config"
	"go-microservice-api/internal/database"
	"go-microservice-api/internal/handlers"
	"go-microservice-api/internal/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.Load()
	auth.InitSecret(cfg.JWTSecret)
	database.Connect(cfg.MongoURI, cfg.DBName)

	r := mux.NewRouter()
	r.Use(middleware.Recovery)
	r.Use(middleware.RequestID)
	r.Use(middleware.CORS(cfg.CORSOrigin))
	r.Use(middleware.Logging)

	// Public routes
	r.HandleFunc("/api/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/api/version", handlers.Version).Methods("GET")
	r.HandleFunc("/api/metrics", handlers.Metrics).Methods("GET")
	r.HandleFunc("/api/auth/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/auth/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/coupons/validate", handlers.ValidateCoupon).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.SimpleRateLimit)
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

	api.HandleFunc("/products/{product_id}/reviews", handlers.ListReviewsForProduct).Methods("GET")
	api.HandleFunc("/products/{product_id}/reviews", handlers.CreateReview).Methods("POST")

	api.HandleFunc("/categories", handlers.ListCategories).Methods("GET")
	api.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	api.HandleFunc("/categories/{id}", handlers.GetCategory).Methods("GET")
	api.HandleFunc("/categories/{id}", handlers.DeleteCategory).Methods("DELETE")

	api.HandleFunc("/cart", handlers.GetCart).Methods("GET")
	api.HandleFunc("/cart/lines", handlers.UpsertCartLine).Methods("PUT")

	api.HandleFunc("/addresses", handlers.ListAddresses).Methods("GET")
	api.HandleFunc("/addresses", handlers.CreateAddress).Methods("POST")
	api.HandleFunc("/addresses/{id}", handlers.DeleteAddress).Methods("DELETE")

	api.HandleFunc("/coupons", handlers.CreateCoupon).Methods("POST")
	api.HandleFunc("/coupons/{id}", handlers.DeleteCoupon).Methods("DELETE")

	api.HandleFunc("/notifications", handlers.ListNotifications).Methods("GET")
	api.HandleFunc("/notifications", handlers.CreateNotification).Methods("POST")
	api.HandleFunc("/notifications/{id}/read", handlers.MarkNotificationRead).Methods("PATCH")

	api.HandleFunc("/wishlist", handlers.ListWishlist).Methods("GET")
	api.HandleFunc("/wishlist", handlers.AddWishlist).Methods("POST")
	api.HandleFunc("/wishlist/{id}", handlers.RemoveWishlist).Methods("DELETE")

	api.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	api.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	api.HandleFunc("/orders/{id}", handlers.GetOrder).Methods("GET")
	api.HandleFunc("/orders/{id}/status", handlers.PatchOrderStatus).Methods("PATCH")
	api.HandleFunc("/orders/{id}/cancel", handlers.CancelOrder).Methods("POST")

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}

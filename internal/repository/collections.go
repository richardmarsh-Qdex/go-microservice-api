package repository

import (
	"go-microservice-api/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func Users() *mongo.Collection {
	return database.DB.Collection("users")
}

func Products() *mongo.Collection {
	return database.DB.Collection("products")
}

func Orders() *mongo.Collection {
	return database.DB.Collection("orders")
}

func Carts() *mongo.Collection {
	return database.DB.Collection("carts")
}

func Coupons() *mongo.Collection {
	return database.DB.Collection("coupons")
}

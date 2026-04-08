package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Rating    int                `json:"rating" bson:"rating"`
	Title     string             `json:"title" bson:"title"`
	Body      string             `json:"body" bson:"body"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type ReviewRequest struct {
	ProductID string `json:"product_id"`
	Rating    int    `json:"rating"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}

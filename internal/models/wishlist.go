package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WishlistItem struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type WishlistAddRequest struct {
	ProductID string `json:"product_id"`
}

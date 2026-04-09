package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartLine struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity  int                `json:"quantity" bson:"quantity"`
}

type Cart struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Lines     []CartLine         `json:"lines" bson:"lines"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CartLineRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

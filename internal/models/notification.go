package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title     string             `json:"title" bson:"title"`
	Message   string             `json:"message" bson:"message"`
	Read      bool               `json:"read" bson:"read"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type NotificationRequest struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

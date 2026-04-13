package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Label      string             `json:"label" bson:"label"`
	Line1      string             `json:"line1" bson:"line1"`
	Line2      string             `json:"line2" bson:"line2"`
	City       string             `json:"city" bson:"city"`
	Region     string             `json:"region" bson:"region"`
	PostalCode string             `json:"postal_code" bson:"postal_code"`
	Country    string             `json:"country" bson:"country"`
	IsDefault  bool               `json:"is_default" bson:"is_default"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
}

type AddressRequest struct {
	Label      string `json:"label"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	City       string `json:"city"`
	Region     string `json:"region"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	IsDefault  bool   `json:"is_default"`
}

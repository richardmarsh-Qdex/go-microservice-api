package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coupon struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code       string             `json:"code" bson:"code"`
	PercentOff float64            `json:"percent_off" bson:"percent_off"`
	MaxUses    int                `json:"max_uses" bson:"max_uses"`
	UsedCount  int                `json:"used_count" bson:"used_count"`
	ValidFrom  time.Time          `json:"valid_from" bson:"valid_from"`
	ValidTo    time.Time          `json:"valid_to" bson:"valid_to"`
	Active     bool               `json:"active" bson:"active"`
}

type CouponRequest struct {
	Code       string    `json:"code"`
	PercentOff float64   `json:"percent_off"`
	MaxUses    int       `json:"max_uses"`
	ValidFrom  time.Time `json:"valid_from"`
	ValidTo    time.Time `json:"valid_to"`
}

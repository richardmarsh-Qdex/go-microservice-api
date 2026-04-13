package services

import (
	"context"
	"errors"
	"time"

	"go-microservice-api/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrOrderNotFound = errors.New("order not found or not owned")

func UpdateOrderStatus(ctx context.Context, orderID primitive.ObjectID, status string) error {
	_, err := database.DB.Collection("orders").UpdateOne(
		ctx,
		bson.M{"_id": orderID},
		bson.M{"$set": bson.M{"status": status, "updated_at": time.Now()}},
	)
	return err
}

func CancelOrder(ctx context.Context, orderID, userID primitive.ObjectID) error {
	res, err := database.DB.Collection("orders").UpdateOne(
		ctx,
		bson.M{"_id": orderID, "user_id": userID},
		bson.M{"$set": bson.M{"status": "cancelled", "updated_at": time.Now()}},
	)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return ErrOrderNotFound
	}
	return nil
}

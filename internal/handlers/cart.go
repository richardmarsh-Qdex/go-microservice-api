package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go-microservice-api/internal/contextkeys"
	"go-microservice-api/internal/database"
	"go-microservice-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getOrCreateCart(ctx context.Context, userOID primitive.ObjectID) (*models.Cart, error) {
	col := database.DB.Collection("carts")
	var c models.Cart
	err := col.FindOne(ctx, bson.M{"user_id": userOID}).Decode(&c)
	if err == nil {
		return &c, nil
	}
	c = models.Cart{
		ID:        primitive.NewObjectID(),
		UserID:    userOID,
		Lines:     nil,
		UpdatedAt: time.Now(),
	}
	_, err = col.InsertOne(ctx, c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	uidStr, ok := r.Context().Value(contextkeys.UserID).(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	userOID, _ := primitive.ObjectIDFromHex(uidStr)
	c, err := getOrCreateCart(r.Context(), userOID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func UpsertCartLine(w http.ResponseWriter, r *http.Request) {
	var req models.CartLineRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uidStr, ok := r.Context().Value(contextkeys.UserID).(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	userOID, _ := primitive.ObjectIDFromHex(uidStr)
	pid, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		http.Error(w, "invalid product_id", http.StatusBadRequest)
		return
	}
	c, err := getOrCreateCart(r.Context(), userOID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	found := false
	for i := range c.Lines {
		if c.Lines[i].ProductID == pid {
			c.Lines[i].Quantity = req.Quantity
			found = true
			break
		}
	}
	if !found {
		c.Lines = append(c.Lines, models.CartLine{ProductID: pid, Quantity: req.Quantity})
	}
	c.UpdatedAt = time.Now()
	_, err = database.DB.Collection("carts").ReplaceOne(
		r.Context(),
		bson.M{"_id": c.ID},
		c,
		options.Replace().SetUpsert(false),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(c)
}

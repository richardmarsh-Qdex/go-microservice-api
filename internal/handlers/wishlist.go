package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"go-microservice-api/internal/contextkeys"
	"go-microservice-api/internal/database"
	"go-microservice-api/internal/httputil"
	"go-microservice-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListWishlist(w http.ResponseWriter, r *http.Request) {
	uidStr, ok := r.Context().Value(contextkeys.UserID).(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	userOID, _ := primitive.ObjectIDFromHex(uidStr)
	cur, err := database.DB.Collection("wishlist").Find(r.Context(), bson.M{"user_id": userOID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(r.Context())
	var out []models.WishlistItem
	if err := cur.All(r.Context(), &out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(out)
}

func AddWishlist(w http.ResponseWriter, r *http.Request) {
	var req models.WishlistAddRequest
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
	item := models.WishlistItem{
		ID:        primitive.NewObjectID(),
		UserID:    userOID,
		ProductID: pid,
		CreatedAt: time.Now(),
	}
	_, err = database.DB.Collection("wishlist").InsertOne(r.Context(), item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func RemoveWishlist(w http.ResponseWriter, r *http.Request) {
	oid, err := httputil.PathID(r, "id")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	uidStr, ok := r.Context().Value(contextkeys.UserID).(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	userOID, _ := primitive.ObjectIDFromHex(uidStr)
	res, err := database.DB.Collection("wishlist").DeleteOne(r.Context(), bson.M{"_id": oid, "user_id": userOID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res.DeletedCount == 0 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

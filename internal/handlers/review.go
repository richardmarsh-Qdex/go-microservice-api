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

func ListReviewsForProduct(w http.ResponseWriter, r *http.Request) {
	oid, err := httputil.PathID(r, "product_id")
	if err != nil {
		http.Error(w, "invalid product_id", http.StatusBadRequest)
		return
	}
	cur, err := database.DB.Collection("reviews").Find(r.Context(), bson.M{"product_id": oid})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(r.Context())
	var out []models.Review
	if err := cur.All(r.Context(), &out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(out)
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	var req models.ReviewRequest
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
	prodOID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		http.Error(w, "invalid product_id", http.StatusBadRequest)
		return
	}
	rev := models.Review{
		ID:        primitive.NewObjectID(),
		ProductID: prodOID,
		UserID:    userOID,
		Rating:    req.Rating,
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
	}
	_, err = database.DB.Collection("reviews").InsertOne(r.Context(), rev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rev)
}

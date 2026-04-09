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

func ListNotifications(w http.ResponseWriter, r *http.Request) {
	uidStr, ok := r.Context().Value(contextkeys.UserID).(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	userOID, _ := primitive.ObjectIDFromHex(uidStr)
	cur, err := database.DB.Collection("notifications").Find(r.Context(), bson.M{"user_id": userOID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(r.Context())
	var out []models.Notification
	if err := cur.All(r.Context(), &out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(out)
}

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	var req models.NotificationRequest
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
	n := models.Notification{
		ID:        primitive.NewObjectID(),
		UserID:    userOID,
		Title:     req.Title,
		Message:   req.Message,
		Read:      false,
		CreatedAt: time.Now(),
	}
	_, err := database.DB.Collection("notifications").InsertOne(r.Context(), n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(n)
}

func MarkNotificationRead(w http.ResponseWriter, r *http.Request) {
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
	_, err = database.DB.Collection("notifications").UpdateOne(
		r.Context(),
		bson.M{"_id": oid, "user_id": userOID},
		bson.M{"$set": bson.M{"read": true}},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

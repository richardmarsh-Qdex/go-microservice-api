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

func ListAddresses(w http.ResponseWriter, r *http.Request) {
	uidStr, ok := r.Context().Value(contextkeys.UserID).(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	userOID, _ := primitive.ObjectIDFromHex(uidStr)
	cur, err := database.DB.Collection("addresses").Find(r.Context(), bson.M{"user_id": userOID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(r.Context())
	var out []models.Address
	if err := cur.All(r.Context(), &out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(out)
}

func CreateAddress(w http.ResponseWriter, r *http.Request) {
	var req models.AddressRequest
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
	a := models.Address{
		ID:         primitive.NewObjectID(),
		UserID:     userOID,
		Label:      req.Label,
		Line1:      req.Line1,
		Line2:      req.Line2,
		City:       req.City,
		Region:     req.Region,
		PostalCode: req.PostalCode,
		Country:    req.Country,
		IsDefault:  req.IsDefault,
		CreatedAt:  time.Now(),
	}
	_, err := database.DB.Collection("addresses").InsertOne(r.Context(), a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}

func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	id, err := httputil.PathID(r, "id")
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
	res, err := database.DB.Collection("addresses").DeleteOne(r.Context(), bson.M{"_id": id, "user_id": userOID})
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

package handlers

import (
	"encoding/json"
	"net/http"

	"go-microservice-api/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	collection := database.DB.Collection("users")
	cursor, err := collection.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	var users []map[string]interface{}
	if err = cursor.All(r.Context(), &users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	collection := database.DB.Collection("users")
	var user map[string]interface{}
	err = collection.FindOne(r.Context(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := database.DB.Collection("users")
	_, err = collection.UpdateOne(
		r.Context(),
		bson.M{"_id": objectID},
		bson.M{"$set": updates},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	collection := database.DB.Collection("users")
	_, err = collection.DeleteOne(r.Context(), bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

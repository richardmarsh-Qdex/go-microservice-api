package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"go-microservice-api/internal/database"
	"go-microservice-api/internal/httputil"
	"go-microservice-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListCategories(w http.ResponseWriter, r *http.Request) {
	col := database.DB.Collection("categories")
	cur, err := col.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(r.Context())
	var out []models.Category
	if err := cur.All(r.Context(), &out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(out)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req models.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c := models.Category{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err := database.DB.Collection("categories").InsertOne(r.Context(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := httputil.PathID(r, "id")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var c models.Category
	err = database.DB.Collection("categories").FindOne(r.Context(), bson.M{"_id": id}).Decode(&c)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := httputil.PathID(r, "id")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	_, err = database.DB.Collection("categories").DeleteOne(r.Context(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"go-microservice-api/internal/models"
	"go-microservice-api/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	collection := database.DB.Collection("products")
	cursor, err := collection.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	var products []models.Product
	if err = cursor.All(r.Context(), &products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	collection := database.DB.Collection("products")
	var product models.Product
	err = collection.FindOne(r.Context(), bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req models.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := models.Product{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		SKU:         req.SKU,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	collection := database.DB.Collection("products")
	_, err := collection.InsertOne(r.Context(), product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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

	updates["updated_at"] = time.Now()

	collection := database.DB.Collection("products")
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

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	collection := database.DB.Collection("products")
	_, err = collection.DeleteOne(r.Context(), bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

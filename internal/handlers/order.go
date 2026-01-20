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

func GetOrders(w http.ResponseWriter, r *http.Request) {
	collection := database.DB.Collection("orders")
	cursor, err := collection.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	var orders []models.Order
	if err = cursor.All(r.Context(), &orders); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	collection := database.DB.Collection("orders")
	var order models.Order
	err = collection.FindOne(r.Context(), bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(string)
	userObjectID, _ := primitive.ObjectIDFromHex(userID)

	var orderItems []models.OrderItem
	var total float64

	productCollection := database.DB.Collection("products")

	for _, itemReq := range req.Items {
		productID, _ := primitive.ObjectIDFromHex(itemReq.ProductID)
		var product models.Product
		err := productCollection.FindOne(r.Context(), bson.M{"_id": productID}).Decode(&product)
		if err != nil {
			http.Error(w, "Product not found", http.StatusBadRequest)
			return
		}

		itemTotal := product.Price * float64(itemReq.Quantity)
		total += itemTotal

		orderItems = append(orderItems, models.OrderItem{
			ProductID: productID,
			Quantity:  itemReq.Quantity,
			Price:     product.Price,
		})
	}

	order := models.Order{
		ID:        primitive.NewObjectID(),
		UserID:    userObjectID,
		Items:     orderItems,
		Total:     total,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	collection := database.DB.Collection("orders")
	_, err := collection.InsertOne(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

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

func ValidateCoupon(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code required", http.StatusBadRequest)
		return
	}
	var c models.Coupon
	err := database.DB.Collection("coupons").FindOne(r.Context(), bson.M{"code": code, "active": true}).Decode(&c)
	if err != nil {
		http.Error(w, "invalid coupon", http.StatusNotFound)
		return
	}
	now := time.Now()
	if now.Before(c.ValidFrom) || now.After(c.ValidTo) {
		http.Error(w, "coupon not valid", http.StatusBadRequest)
		return
	}
	if c.MaxUses > 0 && c.UsedCount >= c.MaxUses {
		http.Error(w, "coupon exhausted", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":        c.Code,
		"percent_off": c.PercentOff,
	})
}

func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	var req models.CouponRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c := models.Coupon{
		ID:         primitive.NewObjectID(),
		Code:       req.Code,
		PercentOff: req.PercentOff,
		MaxUses:    req.MaxUses,
		UsedCount:  0,
		ValidFrom:  req.ValidFrom,
		ValidTo:    req.ValidTo,
		Active:     true,
	}
	_, err := database.DB.Collection("coupons").InsertOne(r.Context(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func DeleteCoupon(w http.ResponseWriter, r *http.Request) {
	id, err := httputil.PathID(r, "id")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	_, err = database.DB.Collection("coupons").DeleteOne(r.Context(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

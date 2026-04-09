package services

import "go-microservice-api/internal/models"

func LineTotal(price float64, qty int) float64 {
	return price * float64(qty)
}

func ApplyPercentDiscount(subtotal, percent float64) float64 {
	if percent <= 0 {
		return subtotal
	}
	return subtotal * (1 - percent/100)
}

func OrderSubtotal(items []models.OrderItem) float64 {
	var t float64
	for _, it := range items {
		t += LineTotal(it.Price, it.Quantity)
	}
	return t
}

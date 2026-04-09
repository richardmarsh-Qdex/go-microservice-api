package httputil

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PathID(r *http.Request, key string) (primitive.ObjectID, error) {
	vars := mux.Vars(r)
	id := vars[key]
	return primitive.ObjectIDFromHex(id)
}

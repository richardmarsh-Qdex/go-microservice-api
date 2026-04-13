package handlers

import (
	"encoding/json"
	"net/http"

	"go-microservice-api/internal/metrics"
	"go-microservice-api/internal/version"
)

func Metrics(w http.ResponseWriter, r *http.Request) {
	out := make(map[string]interface{})
	for k, v := range metrics.Snapshot() {
		out[k] = v
	}
	out["active_connections"] = metrics.ActiveConnections()
	_ = json.NewEncoder(w).Encode(out)
}

func Version(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(version.Info())
}

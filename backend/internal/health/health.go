package health

import (
	"encoding/json"
	"net/http"

	"cloudboard/internal/database"
)

type HealthResponse struct {
	Status string `json:"status"`
}

// HealthHandler returns ok if the HTTP server is alive
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
}

// ReadyHandler returns ok if the app can connect to the database
func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if database.DB == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(HealthResponse{Status: "database_not_initialized"})
		return
	}

	err := database.DB.Ping()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(HealthResponse{Status: "database_unreachable"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Status: "ready"})
}

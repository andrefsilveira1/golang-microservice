package endpoints

import (
	"net/http"

	"microservices/services/internal/domain"
)

func MakeDeleteItemEndpoint(itemService *domain.ItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Endpoint not implemented...", http.StatusInternalServerError)
	}
}

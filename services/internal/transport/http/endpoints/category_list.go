package endpoints

import (
	"net/http"

	"microservices/internal/domain"
)

func MakeListCategoryEndpoint(itemService *domain.ItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "endpoint not implemented yet", http.StatusInternalServerError)
	}
}

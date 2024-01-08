package endpoints

import (
	"net/http"

	"menu-service/internal/domain"
)

func MakeListCategoryEndpoint(ItemService *domain.ItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Endpoint not implemented...", http.StatusInternalServerError)
	}
}

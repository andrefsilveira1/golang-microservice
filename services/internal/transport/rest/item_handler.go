package rest

import (
	"net/http"

	"microservices/services/internal/domain"
	"microservices/services/internal/transport/rest/endpoints"

	"github.com/gorilla/mux"
)

type ItemHandler struct {
	itemService *domain.ItemService
}

func NewItemHandler(itemService *domain.ItemService) *ItemHandler {
	return &ItemHandler{
		itemService: itemService,
	}
}

func (h *ItemHandler) Register(router *mux.Router) {
	listItemEndpoint := endpoints.MakeListItemEndpoint(h.itemService)
	findItemEndpoint := endpoints.MakeGetItemEndpoint(h.itemService)

	router.HandleFunc("/items", listItemEndpoint).Methods(http.MethodGet)
	router.HandleFunc("/items/{id}", findItemEndpoint).Methods(http.MethodGet)

	protected := router.PathPrefix("/").Subrouter()

	createItemEndpoint := endpoints.MakeCreateItemEndpoint(h.itemService)
	updateItemEndpoint := endpoints.MakeUpdateItemEndpoint(h.itemService)
	deleteItemEndpoint := endpoints.MakeDeleteItemEndpoint(h.itemService)

	protected.HandleFunc("/items", createItemEndpoint).Methods(http.MethodPost)
	protected.HandleFunc("/items/{id}", updateItemEndpoint).Methods(http.MethodPut)
	protected.HandleFunc("/items/{id}", deleteItemEndpoint).Methods(http.MethodDelete)
}

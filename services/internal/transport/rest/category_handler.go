package rest

import (
	"microservices/services/internal/domain"
	"microservices/services/internal/transport/rest/endpoints"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	categoryService *domain.CategoryService
}

func NewCategoryHandler(categoryService *domain.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) Register(router *mux.Router) {
	listCategoryEndpoint := endpoints.MakeListCategoryEndpoint(h.categoryService)
	findCategoryEndpoint := endpoints.MakeFindCategoryEndpoint(h.categoryService)

	router.HandleFunc("/categories", listCategoryEndpoint).Methods(http.MethodGet)
	router.HandleFunc("/categories/{id}", findCategoryEndpoint).Methods(http.MethodGet)

	protected := router.PathPrefix("/").Subrouter()

	createCategoryEndpoint := endpoints.MakeCreateCategoryEndpoint(h.categoryService)
	updateCategoryEndpoint := endpoints.MakeUpdateCategoryEndpoint(h.categoryService)
	deleteCategoryEndpoint := endpoints.MakeDeleteCategoryEndpoint(h.categoryService)

	protected.HandleFunc("/categories", createCategoryEndpoint).Methods(http.MethodPost)
	protected.HandleFunc("/categories/{id}", updateCategoryEndpoint).Methods(http.MethodPut)
	protected.HandleFunc("/categories/{id}", deleteCategoryEndpoint).Methods(http.MethodDelete)
}

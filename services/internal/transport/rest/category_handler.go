package rest

import "microservices/services/internal/domain"

type CategoryHandler struct {
	categoryService *domain.CategoryService
}


func NewCategoryHandler(categoryService *domain.CategoryService) *CategoryHandler {
	return &CategoryHandler {
		categoryService: categoryService
	}
}
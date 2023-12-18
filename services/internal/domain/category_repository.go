package domain

type CategoryRepository interface {
	CreateCategory(category *Category) error
	UpdateCategory(category *Category) error
	DeleteCategory(categoryId uint) error
	FindCategoryById(categoryId uint) (*Category, error)
	ListCategories() ([]*Category, error)
}
package domain

type CategoryRepository interface {
	CreateCategory(category *Category) error
	UpdateCategory(category *Category) error
	DeleteCategory(categoryId uint) error
	FindCategoryByID(categoryId uint) (*Category, error)
	ListCategories() ([]*Category, error)
	AddItem(itemId uint, categoryId uint) error
}

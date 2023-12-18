package domain

import ("github.com/pkg/errors")

type CategoryService struct {
	categoryRepository CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService {
		categoryRepository: repo,
	}
}

func (s *CategoryService) CreateCategory(category *Category) error {
	if category == nil {
		return errors.New("Category can not be null")
	}

	return s.categoryRepository.CreateCategory(category)
}

func (s *CategoryService) UpdateCategory(category *Category) error {
	if category == nil {
		return errors.New("Category can not be null")
	}

	return s.categoryRepository.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(category uint) error {
	if category <= 0 {
		return errors.New("Category ID can not be null")
	}

	return s.categoryRepository.DeleteCategory(category)
}

func (s *CategoryService) FindCategoryById(category uint)  (*Category, error) {
	if category <= 0 {
		return nil, errors.New("Category can not be null")
	}

	return s.categoryRepository.FindCategoryById(category)
}

func (s *CategoryService) ListCategories()  ([]*Category, error) {
	return s.categoryRepository.ListCategories()
}



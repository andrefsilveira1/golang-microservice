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


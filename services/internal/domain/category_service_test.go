package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) CreateCategory(category *Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) UpdateCategory(category *Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) DeleteCategory(categoryID uint) error {
	args := m.Called(categoryID)
	return args.Error(0)
}

func (m *MockCategoryRepository) FindCategoryByID(categoryID uint) (*Category, error) {
	args := m.Called(categoryID)
	if category := args.Get(0); category != nil {
		return category.(*Category), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCategoryRepository) ListCategories() ([]*Category, error) {
	args := m.Called()
	return args.Get(0).([]*Category), args.Error(1)
}

func TestCreateCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	category := &Category{Name: "Category Test"}

	repo.On("Create", category).Return(nil)

	err := service.CreateCategory(category)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUpdateCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	category := &Category{Id: 1, Name: "Category Test"}

	repo.On("Update", category).Return(nil)

	err := service.UpdateCategory(category)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	categoryID := uint(1)

	repo.On("Delete", categoryID).Return(nil)

	err := service.DeleteCategory(categoryID)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestGetCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	categoryID := uint(1)
	category := &Category{Id: categoryID, Name: "Category Test"}

	repo.On("Get", categoryID).Return(category, nil)

	result, err := service.FindCategoryByID(categoryID)

	assert.NoError(t, err)
	assert.Equal(t, category, result)
	repo.AssertExpectations(t)
}

func TestListCategories(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	categories := []*Category{
		{Id: 1, Name: "Category 1"},
		{Id: 2, Name: "Category 2"},
	}

	repo.On("List").Return(categories, nil)

	result, err := service.ListCategories()

	assert.NoError(t, err)
	assert.Equal(t, categories, result)
	repo.AssertExpectations(t)
}

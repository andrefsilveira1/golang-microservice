package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockItemRepository struct {
	mock.Mock
}

func (m *MockItemRepository) Create(item *Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockItemRepository) Item(item *Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockItemRepository) Delete(itemID uint) error {
	args := m.Called(itemID)
	return args.Error(0)
}

func (m *MockItemRepository) Get(itemID uint) (*Item, error) {
	args := m.Called(itemID)
	if item := args.Get(0); item != nil {
		return item.(*Item), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockItemRepository) List() ([]*Item, error) {
	args := m.Called()
	return args.Get(0).([]*Item), args.Error(1)
}

func TestCreateItem(t *testing.T) {
	repo := new(MockItemRepository)
	service := NewItemService(repo)

	item := &Item{Name: "Item Test"}

	repo.On("Create", item).Return(nil)

	err := service.Create(item)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUpdateItem(t *testing.T) {
	repo := new(MockItemRepository)
	service := NewItemService(repo)

	item := &Item{ID: 1, Name: "Item Test"}

	repo.On("Update", item).Return(nil)

	err := service.Update(item)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteItem(t *testing.T) {
	repo := new(MockItemRepository)
	service := NewItemService(repo)

	itemID := uint(1)

	repo.On("Delete", itemID).Return(nil)

	err := service.Delete(itemID)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestGetItem(t *testing.T) {
	repo := new(MockItemRepository)
	service := NewItemService(repo)

	itemID := uint(1)
	item := &Item{ID: itemID, Name: "Item Test"}

	repo.On("Get", itemID).Return(item, nil)

	result, err := service.Get(itemID)

	assert.NoError(t, err)
	assert.Equal(t, item, result)
	repo.AssertExpectations(t)
}

func TestListItems(t *testing.T) {
	repo := new(MockItemRepository)
	service := NewItemService(repo)

	items := []*Item{
		{ID: 1, Name: "Item 1"},
		{ID: 2, Name: "Item 2"},
	}

	repo.On("List").Return(items, nil)

	result, err := service.List()

	assert.NoError(t, err)
	assert.Equal(t, items, result)
	repo.AssertExpectations(t)
}

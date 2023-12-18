package domain

type ItemRepository interface {
	CreateItem(item *Item) error
	UpdateItem(item *Item) error
	DeleteItem(itemId uint) error
	FindItemById(itemId uint) (*Item, error)
	ListItems() ([]*Item, error)
}
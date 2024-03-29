package postgres

import (
	"fmt"
	"log"
	"microservices/services/internal/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	createItem = "create item"
	deleteItem = "delete item by id"
	getItem    = "get item by id"
	listItem   = "list item"
	updateItem = "update item by id"
)

type ItemRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func queriesItem() map[string]string {
	return map[string]string{
		createItem: `INSERT INTO items (name, description, price) VALUES ($1, $2, $3) RETURNING *`,
		deleteItem: `UPDATE items SET deleted_at = NOW() WHERE id = $1`,
		getItem:    `SELECT * FROM items WHERE id = $1`,
		listItem:   `SELECT * FROM items WHERE deleted_at IS NULL ORDER BY name ASC`,
		updateItem: `UPDATE items SET name = $1, description = $2, price = $3, updated_at = NOW() WHERE id = $4 RETURNING *`,
	}
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	sqlStatements := make(map[string]*sqlx.Stmt)

	var errs []error
	for queryName, query := range queriesItem() {
		stmt, err := db.Preparex(query)
		if err != nil {
			log.Printf("Error preparing statement %s: %v", queryName, err)
			errs = append(errs, err)
		}

		sqlStatements[queryName] = stmt
	}

	if len(errs) > 0 {
		log.Fatalf("Item repository was not able to build all the statements")
		return nil
	}

	return &ItemRepository{
		DB:         db,
		statements: sqlStatements,
	}
}

func (r *ItemRepository) statement(query string) (*sqlx.Stmt, error) {
	stmt, ready := r.statements[query]
	if !ready {
		return nil, fmt.Errorf("Prepared statement '%s' not found", query)
	}

	return stmt, nil
}

func (r *ItemRepository) CreateItem(item *domain.Item) error {
	stmt, err := r.statement(createItem)
	if err != nil {
		return err
	}

	if err := stmt.Get(item, item.Name, item.Description, item.Price); err != nil {
		return fmt.Errorf("Error while creating new item: %v", err)
	}

	return nil
}

func (r *ItemRepository) UpdateItem(item *domain.Item) error {
	stmt, err := r.statement(updateItem)
	if err != nil {
		return err
	}

	item.UpdatedAt = time.Now()

	params := []interface{}{
		item.Name,
		item.Description,
		item.Price,
		item.Id,
	}

	if err := stmt.Get(item, params...); err != nil {
		return fmt.Errorf("Error until update process")
	}

	return nil
}

func (r *ItemRepository) DeleteItem(itemId uint) error {
	stmt, err := r.statement(deleteItem)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(itemId); err != nil {
		return fmt.Errorf("Error deleting item with id '%d'", itemId)
	}

	return nil
}

func (r *ItemRepository) FindItemById(itemId uint) (*domain.Item, error) {
	stmt, err := r.statement(getItem)
	if err != nil {
		return nil, err
	}

	item := &domain.Item{}
	if err := stmt.Get(item, itemId); err != nil {
		return nil, fmt.Errorf("Error getting the item with id '%d' ", itemId)
	}

	return item, nil
}

func (r *ItemRepository) ListItems() ([]*domain.Item, error) {
	stmt, err := r.statement(listItem)
	if err != nil {
		return nil, err
	}

	var items []*domain.Item
	if err := stmt.Select(&items); err != nil {
		return nil, fmt.Errorf("Error listing items")
	}

	return items, nil
}

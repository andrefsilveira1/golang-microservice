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

func (r *ItemRepository) Create(item *domain.Item) error {
	stmt, err := r.statement(createItem)
	if err != nil {
		return err
	}

	if err := stmt.Get(item, item.name, item.description, item.price); err != nil {
		return fmt.Errorf("Error while creating new item: %v", err)
	}

	return nil
}

func (r *ItemRepository) Update(item *domain.Item) error {
	stmt, err := r.statement(updateItem)
	if err != nil {
		return err
	}

	item.updated_at = time.Now()

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

func (r *ItemRepository) Delete(itemId int) error {
	stmt, err := r.statement(deleteItem)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(itemId); err != nil {
		return fmt.Errorf("Error deleting item with id '%d'", itemId)
	}

	return nil
}

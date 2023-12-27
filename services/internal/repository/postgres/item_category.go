package postgres

import (
	"log"

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
}

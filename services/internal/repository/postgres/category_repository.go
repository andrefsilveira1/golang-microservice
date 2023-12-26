package postgres


import (
	"fmt"
)

const (
	createCategory = "create category"
	deleteCategory = "delete category by id"
	getCategory    = "get category by id"
	listCategory   = "list category"
	updateCategory = "update category by id"
)

func queriesCategory() map[string]string {
	return map[string]string{
		createCategory: `INSERT INTO categories (name) VALUES ($1) RETURNING *`,
		deleteCategory: `UPDATE categories SET deleted_at = NOW() WHERE id = $1`,
		getCategory:    `SELECT * FROM categories WHERE id = $1`,
		listCategory:   `SELECT * FROM categories WHERE deleted_at IS NULL BY name ASC`,
		updateCategory: `UPDATE categories SET name = $1, updated_at = NOW() WHERE id = $2 RETURNING *`,
	}
}
package postgres

import (
	"context"
	"fmt"
	"log"
	"microservices/services/internal/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	createCategory    = "create category"
	deleteCategory    = "delete category by id"
	getCategory       = "get category by id"
	listCategory      = "list category"
	updateCategory    = "update category by id"
	addItem           = "add item to category"
	removeItem        = "remove item from category"
	getItemCategories = "get item categories"
)

type CategoryRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func queriesCategory() map[string]string {
	return map[string]string{
		createCategory: `INSERT INTO categories (name) VALUES ($1) RETURNING *`,
		deleteCategory: `UPDATE categories SET deleted_at = NOW() WHERE id = $1`,
		getCategory:    `SELECT * FROM categories WHERE id = $1`,
		listCategory:   `SELECT * FROM categories WHERE deleted_at IS NULL BY name ASC`,
		updateCategory: `UPDATE categories SET name = $1, updated_at = NOW() WHERE id = $2 RETURNING *`,
		addItem:        `INSERT INTO item_categories (item_id, category_id) VALUES (:item_id, :category_id)`,
		removeItem:     `DELETE FROM item_categories WHERE item_id = :item_id AND category_id = :category_id`,
	}
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	statements := make(map[string]*sqlx.Stmt)

	var errors []error
	for queryName, query := range queriesCategory() {
		stmt, err := db.Preparex(query)
		if err != nil {
			log.Printf("Errror preparing statement %s: %v", queryName, err)
			errors = append(errors, err)
		}
		statements[queryName] = stmt
	}

	if len(errors) > 0 {
		log.Fatalf("Category repository was not able to build all the statements")
		return nil
	}

	return &CategoryRepository{
		DB:         db,
		statements: statements,
	}
}

func (r *CategoryRepository) statement(query string) (*sqlx.Stmt, error) {
	stmt, ready := r.statements[query]
	if !ready {
		return nil, fmt.Errorf("Prepared statement '%s' not found", query)
	}

	return stmt, nil
}

func (r *CategoryRepository) CreateCategory(category *domain.Category) error {
	stmt, err := r.statement(createCategory)
	if err != nil {
		return err
	}

	if err := stmt.Get(category, category.Name); err != nil {
		if isUniqueViolationError(err) {
			return fmt.Errorf("Category with name '%s' already exists", category.Name)
		}
		return fmt.Errorf("Error creating category: %v", err)
	}

	return nil
}

func (r *CategoryRepository) UpdateCategory(category *domain.Category) error {
	stmt, err := r.statement(updateCategory)
	if err != nil {
		return err
	}

	category.UpdatedAt = time.Now()

	params := []interface{}{
		category.Name,
		category.Id,
	}

	if err := stmt.Get(category, params...); err != nil {
		// TO implement
		if isUniqueViolationError(err) {
			return fmt.Errorf("Category with name '%s' already exists", category.Name)
		}

		return fmt.Errorf("Error updating category")
	}

	return nil
}

func (r *CategoryRepository) DeleteCategory(categoryId uint) error {
	stmt, err := r.statement(deleteCategory)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(categoryId); err != nil {
		return fmt.Errorf("Error deleting category with id '%d' ", categoryId)
	}
	return nil
}

func (r *CategoryRepository) FindCategoryByID(categoryId uint) (*domain.Category, error) {
	stmt, err := r.statement(getCategory)
	if err != nil {
		return nil, err
	}

	category := &domain.Category{}
	if err := stmt.Get(category, categoryId); err != nil {
		return nil, fmt.Errorf("Error getting the category with id '%d'", categoryId)
	}

	return category, nil

}

func (r *CategoryRepository) ListCategories() ([]*domain.Category, error) {
	stmt, err := r.statement(listCategory)
	if err != nil {
		return nil, err
	}

	var categories []*domain.Category
	if err := stmt.Select(&categories); err != nil {
		return nil, fmt.Errorf("Error getting categories")
	}

	return categories, nil
}

func (r *CategoryRepository) AddItem(itemID uint, categoryID uint) error {
	stmt, err := r.statement(addItem)
	if err != nil {
		return err
	}

	params := []interface{}{
		itemID,
		categoryID,
	}

	if _, err := stmt.ExecContext(context.Background(), params); err != nil {
		return fmt.Errorf("error adding item '%d' from to category '%d'", itemID, categoryID)
	}

	return nil
}

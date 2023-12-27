package postgres

const (
	createItem = "create item"
	deleteItem = "delete item by id"
	getItem    = "get item by id"
	listItem   = "list item"
	updateItem = "update item by id"
)

func queriesItem() map[string]string {
	return map[string]string{
		createItem: `INSERT INTO items (name, description, price) VALUES ($1, $2, $3) RETURNING *`,
		deleteItem: `UPDATE items SET deleted_at = NOW() WHERE id = $1`,
		getItem:    `SELECT * FROM items WHERE id = $1`,
		listItem:   `SELECT * FROM items WHERE deleted_at IS NULL ORDER BY name ASC`,
		updateItem: `UPDATE items SET name = $1, description = $2, price = $3, updated_at = NOW() WHERE id = $4 RETURNING *`,
	}
}

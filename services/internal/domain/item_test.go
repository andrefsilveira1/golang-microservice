package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemJSONSerialization(t *testing.T) {
	item := &Item{
		Id:          1,
		Name:        "Test Item",
		Description: "Description of test item",
		Price:       9.99,
		Categories: []*Category{
			&Category{Name: "Category 1"},
			&Category{Name: "category 2"},
		},
	}

	data, err := json.Marshal(item)
	assert.NoError(t, err, "Error serializing item to JSON")

	var newItem Item
	err = json.Unmarshal(data, &newItem)
	assert.NoError(t, err, "Error deserializing JSON to item")
	assert.Equal(t, item, &newItem, "Original item and deserialized item do not match")
}
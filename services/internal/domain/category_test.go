package domain

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCategoryJSON(t *testing.T) {
	category := &Category{
		Id: 2,
		Name: "Testing",
	}

	data, err:= json.Marshal(category)
	assert.NoError(t, err, "Error: serializing Category (JSON)")

	var tempCategory Category
	err = json.Unmarshal(data, &tempCategory)
	assert.NoError(t, err, "Error: deserializing JSON to Category")

	assert.Equal(t, category, &tempCategory, "The categories are not Equal")
}
// internal/transport/rest/endpoints/category_delete_test.go

package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"microservices/services/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestMakeDeleteCategoryEndpoint(t *testing.T) {
	categoryService := &domain.CategoryService{}

	server := httptest.NewServer(MakeDeleteCategoryEndpoint(categoryService))
	defer server.Close()

	req, err := http.NewRequest("DELETE", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

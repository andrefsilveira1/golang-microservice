package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"microservices/services/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestMakeListCategoryEndpoint(t *testing.T) {
	categoryService := &domain.CategoryService{}

	server := httptest.NewServer(MakeListCategoryEndpoint(categoryService))
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL, nil)
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

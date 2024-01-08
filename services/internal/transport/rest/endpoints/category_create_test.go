package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"microservices/services/internal/domain"
)

func TestMakeCreateCategoryEndpoint(t *testing.T) {
	categoryService := &domain.CategoryService{}

	server := httptest.NewServer(MakeCreateCategoryEndpoint(categoryService))
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL, nil)
	if err != nil {
		t.Fatalf("Error creating HTTP request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, but got %d", resp.StatusCode)
	}
}

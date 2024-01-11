package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"microservices/services/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestMakeUpdateItemEndpoint(t *testing.T) {
	itemService := &domain.ItemService{}
	server := httptest.NewServer(MakeUpdateItemEndpoint(itemService))
	defer server.Close()
	req, err := http.NewRequest("PUT", server.URL, nil)
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

package handlers_test

import (
	"encoding/json"
	"github.com/paveldroo/tdd-in-go-book/bookswap/db"
	"github.com/paveldroo/tdd-in-go-book/bookswap/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIndexIntegration(t *testing.T) {
	if os.Getenv("LONG") == "" {
		t.Skip("Skipping TestIndexIntegration in short mode")
	}
	// Arrange
	b := db.Book{
		ID:      "123",
		Name:    "A new book for testing",
		Author:  "Some testing author",
		OwnerID: "some_owner",
		Status:  db.Available.String(),
	}
	bs := db.NewBookService([]db.Book{b}, nil)
	h := handlers.NewHandler(bs, nil)

	svr := httptest.NewServer(http.HandlerFunc(h.Index))
	defer svr.Close()
	// Act
	r, err := http.Get(svr.URL)
	// Assert
	require.Nil(t, err)
	assert.Equal(t, r.StatusCode, http.StatusOK)
	body, err := io.ReadAll(r.Body)
	require.Nil(t, err)
	_ = r.Body.Close()
	var resp handlers.Response
	err = json.Unmarshal(body, &resp)
	require.Nil(t, err)
	assert.Equal(t, 1, len(resp.Books))
	assert.Contains(t, resp.Books, b)
}

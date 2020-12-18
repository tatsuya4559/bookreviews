package bookreviews

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tatsuya4559/bookreviews/models"
	"go.uber.org/dig"
)

func TestBooksIndex(t *testing.T) {
	container = dig.New()
	container.Provide(mock) //DIコンテナにモックを登録

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	booksIndex(rec, req)

	expected := "123, hoge, hogetaro\n"
	if body := rec.Body.String(); body != expected {
		t.Errorf("booksIndex response = %q, want %q", body, expected)
	}
}

func mock() BookRepository {
	store := make(map[uint]*models.Book)
	store[123] = &models.Book{
		ISBN:      123,
		Title:     "hoge",
		Publisher: "hogetaro",
	}
	return &mockBookRepository{store}
}

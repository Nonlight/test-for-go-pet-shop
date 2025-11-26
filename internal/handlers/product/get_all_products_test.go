package product

import (
	"errors"
	"go-pet-shop/internal/models"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllProducts_Success(t *testing.T) {
	mock := &ProductsMock{
		GetAllProductsFunc: func() ([]models.Product, error) {
			return []models.Product{
				{ID: 1, Name: "Dog Food"},
			}, nil
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	handler := GetAllProducts(slog.Default(), mock)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestGetAllProducts_Error(t *testing.T) {
	mock := &ProductsMock{
		GetAllProductsFunc: func() ([]models.Product, error) {
			return nil, errors.New("DB error")
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	handler := GetAllProducts(slog.Default(), mock)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

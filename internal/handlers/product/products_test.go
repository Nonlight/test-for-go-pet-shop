package product

import (
	"context"
	"errors"
	"go-pet-shop/internal/models"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

// Get Product - Ready
func TestGetAllProducts_Success(t *testing.T) {
	// Мокаем storage — он вернёт один продукт.
	mock := &ProductsMock{
		GetAllProductsFunc: func(ctx context.Context) ([]models.Product, error) {
			return []models.Product{
				{ID: 1, Name: "Dog Food"},
			}, nil
		},
	}

	// Создаем HTTP-запрос GET /products
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	// Создаем хендлер с мок-хранилищем
	handler := New(slog.Default(), mock)

	// Вызываем метод GetAllProducts, который является http.HandlerFunc
	handler.GetAllProducts(w, req)

	// Проверяем HTTP-код
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}
func TestGetAllProducts_Error(t *testing.T) {
	// Мокаем storage — он будет возвращать ошибку
	mock := &ProductsMock{
		GetAllProductsFunc: func(ctx context.Context) ([]models.Product, error) {
			return nil, errors.New("DB error")
		},
	}

	// Создаем запрос
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	handler := New(slog.Default(), mock)
	handler.GetAllProducts(w, req)

	// Ожидаем HTTP 500
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

// =======================
// Create Product
// =======================

func TestCreateProduct_Success(t *testing.T) {
	mock := &ProductsMock{
		CreateProductFunc: func(ctx context.Context, product models.Product) (int, error) {
			return 1, nil
		},
	}

	body := `{"name":"Cat Food","price":12.5,"stock":10}`

	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
	w := httptest.NewRecorder()

	handler := New(slog.Default(), mock)

	handler.CreateProduct(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

}

func TestCreateProduct_BadRequest(t *testing.T) {
	mock := &ProductsMock{}

	body := `{"name":"Cat Food","price":12.5,"stock":10`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
	w := httptest.NewRecorder()

	handler := New(slog.Default(), mock)

	handler.CreateProduct(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
}

func TestCreateProduct_Fail(t *testing.T) {
	mock := &ProductsMock{
		CreateProductFunc: func(ctx context.Context, product models.Product) (int, error) {
			return 0, errors.New("DB error")
		},
	}

	body := `{"name":"Cat Food","price":12.5,"stock":10}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
	w := httptest.NewRecorder()

	handler := New(slog.Default(), mock)
	handler.CreateProduct(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

// =======================
// Update Product
// =======================

func TestUpdateProduct_Success(t *testing.T) {
	mock := &ProductsMock{
		UpdateProductFunc: func(ctx context.Context, product models.Product) error {
			if product.ID != 1 {
				t.Fatalf("expected product ID 1, got %d", product.ID)
			}
			return nil
		},
	}
	body := `{"name":"Cat Food","price":15.5,"stock":20}`
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req := httptest.NewRequest(http.MethodPut, "/products/1", strings.NewReader(body))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler := New(slog.Default(), mock)
	handler.UpdateProduct(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestUpdateProduct_BadRequest(t *testing.T) {
	mock := &ProductsMock{}

	body := `{"name":"Cat Food","price":15.5,"stock":20`
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req := httptest.NewRequest(http.MethodPut, "/products/1", strings.NewReader(body))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	handler := New(slog.Default(), mock)
	handler.UpdateProduct(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
}

func TestUpdateProduct_Fail(t *testing.T) {
	mock := &ProductsMock{
		UpdateProductFunc: func(ctx context.Context, product models.Product) error {
			return errors.New("DB error")
		},
	}

	body := `{"name":"Cat Food","price":15.5,"stock":20}`
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req := httptest.NewRequest(http.MethodPut, "/products/1", strings.NewReader(body))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler := New(slog.Default(), mock)
	handler.UpdateProduct(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

// =======================
// Delete Product
// =======================

func TestDeleteProduct_Success(t *testing.T) {
	mock := &ProductsMock{
		DeleteProductFunc: func(ctx context.Context, id int) error {
			if id != 1 {
				t.Fatalf("expected id 1, got %d", id)
			}
			return nil
		},
	}
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	handler := New(slog.Default(), mock)

	handler.DeleteProduct(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestDeleteProduct_BadRequest(t *testing.T) {
	mock := &ProductsMock{}

	rctx := chi.NewRouteContext()

	req := httptest.NewRequest(http.MethodDelete, "/products/", nil)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	handler := New(slog.Default(), mock)
	handler.DeleteProduct(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
}

func TestDeleteProduct_Fail(t *testing.T) {
	mock := &ProductsMock{
		DeleteProductFunc: func(ctx context.Context, id int) error {
			return errors.New("DB error")
		},
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	handler := New(slog.Default(), mock)

	handler.DeleteProduct(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

package routes

import (
	"kasir-api/handlers"
	"net/http"
)

// SetupCategoryRoutes mengatur semua route untuk category
func SetupCategoryRoutes(categoryHandler *handlers.CategoryHandler) {
	// GET /api/category - Get all categories
	// POST /api/category - Create new category
	http.HandleFunc("/api/category", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetAll(w, r)
		case http.MethodPost:
			categoryHandler.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// GET /api/category/{id} - Get category by ID
	// PUT /api/category/{id} - Update category
	// DELETE /api/category/{id} - Delete category
	http.HandleFunc("/api/category/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetByID(w, r)
		case http.MethodPut:
			categoryHandler.Update(w, r)
		case http.MethodDelete:
			categoryHandler.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

package routes

import (
	"database/sql"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"kasir-api/utils"
	"net/http"
)

// SetupRoutes mengatur semua route aplikasi termasuk dependency injection
func SetupRoutes(db *sql.DB) {
	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		utils.SuccessResponse(w, "API is running", nil, http.StatusOK)
	})

	// Dependency injection untuk Product routes
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)
	SetupProductRoutes(productHandler)

	// Dependency injection untuk Category routes
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	SetupCategoryRoutes(categoryHandler)
}

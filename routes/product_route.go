package routes

import (
	"kasir-api/handlers"
	"kasir-api/middlewares"
	"net/http"
)

// SetupProductRoutes mengatur semua route untuk product
func SetupProductRoutes(productHandler *handlers.ProductHandler, apiKeyMiddleware func(http.HandlerFunc) http.HandlerFunc) {
	// GET /api/product - Get all products
	// POST /api/product - Create new product
	http.HandleFunc("/api/product", middlewares.Logger(middlewares.CORS(apiKeyMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetAll(w, r)
		case http.MethodPost:
			productHandler.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))))

	// GET /api/product/{id} - Get product by ID
	// PUT /api/product/{id} - Update product
	// DELETE /api/product/{id} - Delete product
	http.HandleFunc("/api/product/", middlewares.Logger(middlewares.CORS(apiKeyMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetByID(w, r)
		case http.MethodPut:
			productHandler.Update(w, r)
		case http.MethodDelete:
			productHandler.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))))
}

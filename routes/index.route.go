package routes

import (
	"kasir-api/utils"
	"net/http"
)


func SetupRoutes() {
	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.SuccessResponse(w, "API is running", nil, http.StatusOK)
	})

	// Product routes
	ProductRoutes()

	// Category routes
	CategoryRoutes()
}

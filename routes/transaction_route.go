package routes

import (
	"kasir-api/handlers"
	"kasir-api/middlewares"
	"kasir-api/utils"
	"net/http"
)

func SetupTransactionRoutes(transactionHandler *handlers.TransactionHandler, apiKeyMiddleware func(http.HandlerFunc) http.HandlerFunc) {
	// POST /api/transaction/checkout - Checkout transaction
	http.HandleFunc("/api/transaction/checkout", middlewares.Logger(middlewares.CORS(apiKeyMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			transactionHandler.Checkout(w, r)
		default:
			utils.ErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))))
}

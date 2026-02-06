package routes

import (
	"kasir-api/handlers"
	"kasir-api/utils"
	"net/http"
)

func SetupTransactionRoutes(transactionHandler *handlers.TransactionHandler) {
	// POST /api/transaction/checkout - Checkout transaction
	http.HandleFunc("/api/transaction/checkout", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			transactionHandler.Checkout(w, r)
		default:
			utils.ErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

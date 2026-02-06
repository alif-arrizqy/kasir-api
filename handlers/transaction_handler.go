package handlers

import (
	"encoding/json"
	"fmt"
	"kasir-api/models"
	"kasir-api/services"
	"kasir-api/utils"
	"net/http"
	"strings"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate quantities
	for _, it := range req.Items {
		if it.Quantity <= 0 {
			utils.ErrorResponse(w, fmt.Sprintf("invalid quantity for product with id %d", it.ProductID), http.StatusBadRequest)
			return
		}
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		// if repo returned product not found error, map to 404
		if strings.Contains(err.Error(), "product with id") && strings.Contains(err.Error(), "not found") {
			utils.ErrorResponse(w, err.Error(), http.StatusNotFound)
			return
		}

		utils.ErrorResponse(w, "Failed to process checkout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

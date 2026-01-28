package handlers

import (
	"encoding/json"
	"kasir-api/helper"
	"kasir-api/models"
	"kasir-api/services"
	"kasir-api/utils"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// GetAll - GET /api/product
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, "Products retrieved successfully", products, http.StatusOK)
}

// Create - POST /api/product
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate product
	errorMessage, isValid := helper.ValidateProduct(product)
	if !isValid {
		utils.ErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, "Product created successfully", product, http.StatusCreated)
}

// GetByID - GET /api/product/{id}
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.SuccessResponse(w, "Product retrieved successfully", product, http.StatusOK)
}

// Update - PUT /api/product/{id}
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.SuccessResponse(w, "Product updated successfully", product, http.StatusOK)
}

// Delete - DELETE /api/product/{id}
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, "Product deleted successfully", nil, http.StatusOK)
}
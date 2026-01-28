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

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// GetAll - GET /api/category
func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, "Categories retrieved successfully", categories, http.StatusOK)
}

// Create - POST /api/category
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate category
	errorMessage, isValid := helper.ValidateCategory(category)
	if !isValid {
		utils.ErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	err = h.service.Create(&category)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, "Category created successfully", category, http.StatusCreated)
}

// GetByID - GET /api/category/{id}
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.SuccessResponse(w, "Category retrieved successfully", category, http.StatusOK)
}

// Update - PUT /api/category/{id}
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	category.ID = id
	err = h.service.Update(&category)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.SuccessResponse(w, "Category updated successfully", category, http.StatusOK)
}

// Delete - DELETE /api/category/{id}
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, "Category deleted successfully", nil, http.StatusOK)
}
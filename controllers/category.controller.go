package controllers

import (
	"encoding/json"
	"kasir-api/helper"
	"kasir-api/services"
	"kasir-api/utils"
	"net/http"
	"strconv"
	"strings"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	products := services.GetAllCategories()
	utils.SuccessResponse(w, "Categories retrieved successfully", products, http.StatusOK)
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	// parse id dari url: /api/categories/{id}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	category, err := services.GetCategoryById(id)
	if err != nil {
		utils.ErrorResponse(w, "Category not found", http.StatusNotFound)
		return
	}

	utils.SuccessResponse(w, "Category retrieved successfully", category, http.StatusOK)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// baca data dari request
	var newCategory utils.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)

	if err != nil {
		utils.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// validasi field category
	errorMessage, isValid := helper.ValidateCategory(newCategory)
	if !isValid {
		utils.ErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	category := services.CreateCategory(newCategory)
	utils.SuccessResponse(w, "Category created successfully", category, http.StatusCreated)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari url
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// get data dari request body
	var updateCategory utils.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// validasi field category (setelah decode)
	errorMessage, isValid := helper.ValidateCategory(updateCategory)
	if !isValid {
		utils.ErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	category, err := services.UpdateCategory(id, updateCategory)
	if err != nil {
		utils.ErrorResponse(w, "Category not found", http.StatusNotFound)
	}
	
	utils.SuccessResponse(w, "Category updated successfully", category, http.StatusOK)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari url
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteCategory(id)
	if err != nil {
		utils.ErrorResponse(w, "Category not found", http.StatusNotFound)
	}
	utils.SuccessResponse(w, "Category deleted successfully", nil, http.StatusOK)
}

package services

import (
	"encoding/json"
	"kasir-api/utils"
	"kasir-api/helper"
	"net/http"
	"strconv"
	"strings"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	utils.SuccessResponse(w, "Categories retrieved successfully", utils.Categories, http.StatusOK)
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	// parse id dari url: /api/categories/{id}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// cari kategori dengan id yang sesuai
	for _, c := range utils.Categories {
		if c.ID == id {
			utils.SuccessResponse(w, "Category retrieved successfully", c, http.StatusOK)
			return
		}
	}

	// jika kategori tidak ditemukan
	utils.ErrorResponse(w, "Category not found", http.StatusNotFound)
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

	// masukin data ke dalam array kategori
	newCategory.ID = len(utils.Categories) + 1
	utils.Categories = append(utils.Categories, newCategory)

	utils.SuccessResponse(w, "Category created successfully", newCategory, http.StatusCreated)
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

	// loop categories, cari id yang sesuai, update data
	for i := range utils.Categories {
		if utils.Categories[i].ID == id {
			updateCategory.ID = id // pastikan ID tetap sama
			utils.Categories[i] = updateCategory

			utils.SuccessResponse(w, "Category updated successfully", updateCategory, http.StatusOK)
			return
		}
	}

	// jika kategori tidak ditemukan
	utils.ErrorResponse(w, "Category not found", http.StatusNotFound)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari url
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// loop categories, cari ID, hapus dari slice
	for i, _ := range utils.Categories {
		if utils.Categories[i].ID == id {
			// hapus dengan append slice sebelum dan sesudah index
			utils.Categories = append(utils.Categories[:i], utils.Categories[i+1:]...)

			utils.SuccessResponse(w, "Category deleted successfully", nil, http.StatusOK)
			return
		}
	}

	// jika kategori tidak ditemukan
	utils.ErrorResponse(w, "Category not found", http.StatusNotFound)
}

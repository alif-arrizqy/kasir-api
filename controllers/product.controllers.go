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

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := services.GetAllProducts()
	utils.SuccessResponse(w, "Products retrieved successfully", products, http.StatusOK)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	// parse id dari url
	// url: /api/products/{id}

	// 	### Penjelasan Detail
	// **URL Path Parsing** - Ini yang paling tricky:

	// 1. **URL:** `/api/products/123`
	// 2. **TrimPrefix:** Hilangkan `/api/products/` → dapat `"123"`
	// 3. **Atoi:** Convert `"123"` string → `123` integer

	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	product, err := services.GetProductById(id)
	if err != nil {
		utils.ErrorResponse(w, "Product not found", http.StatusNotFound)
		return
	}

	utils.SuccessResponse(w, "Product retrieved successfully", product, http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// baca data dari request
	var newProduct utils.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// validasi field product
	errorMessage, isValid := helper.ValidateProduct(newProduct)
	if !isValid {
		utils.ErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	product := services.CreateProduct(newProduct)
	utils.SuccessResponse(w, "Product created successfully", product, http.StatusCreated)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")

	// convert id to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// get data dari request body
	var updateProduct utils.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// validasi field product (setelah decode)
	errorMessage, isValid := helper.ValidateProduct(updateProduct)
	if !isValid {
		utils.ErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	product, err := services.UpdateProduct(id, updateProduct)
	if err != nil {
		utils.ErrorResponse(w, "Product not found", http.StatusNotFound)
		return
	}

	utils.SuccessResponse(w, "Product updated successfully", product, http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteProduct(id)
	if err != nil {
		utils.ErrorResponse(w, "Product not found", http.StatusNotFound)
		return
	}

	utils.SuccessResponse(w, "Product deleted successfully", nil, http.StatusOK)
}
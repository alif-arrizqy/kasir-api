package services

import (
	"encoding/json"
	"kasir-api/utils"
	"kasir-api/helper"
	"net/http"
	"strconv"
	"strings"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	utils.SuccessResponse(w, "Products retrieved successfully", utils.Products, http.StatusOK)
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

	// insert data ke dalam array produk
	newProduct.ID = len(utils.Products) + 1       // generate id auto increment
	utils.Products = append(utils.Products, newProduct) // insert data ke dalam array produk

	utils.SuccessResponse(w, "Product created successfully", newProduct, http.StatusCreated)
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

	// cari produk dengan id yang sesuai
	for _, p := range utils.Products {
		if p.ID == id {
			utils.SuccessResponse(w, "Product retrieved successfully", p, http.StatusOK)
			return
		}
	}

	// jika produk tidak ditemukan
	utils.ErrorResponse(w, "Product not found", http.StatusNotFound)
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

	// loop produk, cari id yang sesuai, ganti data sesuai dengan data dari request
	for i := range utils.Products {
		if utils.Products[i].ID == id {
			updateProduct.ID = id
			utils.Products[i] = updateProduct

			utils.SuccessResponse(w, "Product updated successfully", updateProduct, http.StatusOK)
			return
		}
	}

	// jika produk tidak ditemukan
	utils.ErrorResponse(w, "Product not found", http.StatusNotFound)
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

	// loop produk cari ID, dapet index yang mau dihapus
	for i, p := range utils.Products {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			utils.Products = append(utils.Products[:i], utils.Products[i+1:]...)

			utils.SuccessResponse(w, "Product deleted successfully", nil, http.StatusOK)
			return
		}
	}

	// jika produk tidak ditemukan
	utils.ErrorResponse(w, "Product not found", http.StatusNotFound)
}

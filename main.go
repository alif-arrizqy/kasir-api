package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// blueprint untuk produk
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Response wrapper untuk konsistensi
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// in-memory storage
var products = []Product{
	{ID: 1, Name: "Indomie", Price: 7000, Stock: 10},
	{ID: 2, Name: "Vit 100ml", Price: 1000, Stock: 20},
	{ID: 3, Name: "Kecap", Price: 3000, Stock: 20},
}

var categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Makanan ringan"},
	{ID: 2, Name: "Minuman", Description: "Minuman segar"},
	{ID: 3, Name: "Lainnya", Description: "Lainnya"},
}

// Helper untuk success response
func successResponse(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Helper untuk error response
func errorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: message,
		Data:    nil,
	})
}

// Validation Handler
func validateProduct(p Product) (string, bool) {
	// Validasi Name
	if p.Name == "" {
		return "Field 'name' is required", false
	}

	// Validasi Price
	if p.Price <= 0 {
		return "Field 'price' is required and must be greater than 0", false
	}

	// Validasi Stock
	if p.Stock <= 0 {
		return "Field 'stock' is required and must be greater than 0", false
	}

	// Semua validasi passed
	return "", true
}

// Validasi Category - return (errorMessage, isValid)
func validateCategory(c Category) (string, bool) {
	// Validasi Name
	if c.Name == "" {
		return "Field 'name' is required", false
	}

	// Semua validasi passed
	return "", true
}

// Product Handler
func getAllProducts(w http.ResponseWriter, r *http.Request) {
	successResponse(w, "Products retrieved successfully", products, http.StatusOK)
}

// POST localhost:8889/api/products
func createProduct(w http.ResponseWriter, r *http.Request) {
	// baca data dari request
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		errorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// validasi field product
	errorMessage, isValid := validateProduct(newProduct)
	if !isValid {
		errorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	// insert data ke dalam array produk
	newProduct.ID = len(products) + 1       // generate id auto increment
	products = append(products, newProduct) // insert data ke dalam array produk

	successResponse(w, "Product created successfully", newProduct, http.StatusCreated)
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	// parse id dari url
	// url: /api/produk/{id}

	// 	### Penjelasan Detail
	// **URL Path Parsing** - Ini yang paling tricky:

	// 1. **URL:** `/api/produk/123`
	// 2. **TrimPrefix:** Hilangkan `/api/produk/` → dapat `"123"`
	// 3. **Atoi:** Convert `"123"` string → `123` integer

	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
	}

	// cari produk dengan id yang sesuai
	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// jika produk tidak ditemukan
	json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Product not found"})
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// convert id to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateProduct Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	// loop produk, cari id yang sesuai, ganti data sesuai dengan data dari request
	for i := range produk {
		if produk[i].ID == id {
			updateProduct.ID = id
			produk[i] = updateProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduct)
			return
		}
	}

	// jika produk tidak ditemukan
	json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Product not found"})
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// loop produk cari ID, dapet index yang mau dihapus
	for i, p := range produk {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			products = append(products[:i], products[i+1:]...)

			successResponse(w, "Product deleted successfully", nil, http.StatusOK)
			return
		}
	}

	// jika produk tidak ditemukan
	errorResponse(w, "Product not found", http.StatusNotFound)
}

// Category Handler
func getAllCategories(w http.ResponseWriter, r *http.Request) {
	successResponse(w, "Categories retrieved successfully", categories, http.StatusOK)
}

func getCategoryById(w http.ResponseWriter, r *http.Request) {
	// parse id dari url: /api/categories/{id}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// cari kategori dengan id yang sesuai
	for _, c := range categories {
		if c.ID == id {
			successResponse(w, "Category retrieved successfully", c, http.StatusOK)
			return
		}
	}

	// jika kategori tidak ditemukan
	errorResponse(w, "Category not found", http.StatusNotFound)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	// baca data dari request
	var newCategory Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		errorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// validasi field category
	errorMessage, isValid := validateCategory(newCategory)
	if !isValid {
		errorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	// masukin data ke dalam array kategori
	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)

	successResponse(w, "Category created successfully", newCategory, http.StatusCreated)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari url
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// get data dari request body
	var updateCategory Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		errorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// validasi field category (setelah decode)
	errorMessage, isValid := validateCategory(updateCategory)
	if !isValid {
		errorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	// loop categories, cari id yang sesuai, update data
	for i := range categories {
		if categories[i].ID == id {
			updateCategory.ID = id // pastikan ID tetap sama
			categories[i] = updateCategory

			successResponse(w, "Category updated successfully", updateCategory, http.StatusOK)
			return
		}
	}

	// jika kategori tidak ditemukan
	errorResponse(w, "Category not found", http.StatusNotFound)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari url
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// loop categories, cari ID, hapus dari slice
	for i, c := range categories {
		if c.ID == id {
			// hapus dengan append slice sebelum dan sesudah index
			categories = append(categories[:i], categories[i+1:]...)

			successResponse(w, "Category deleted successfully", nil, http.StatusOK)
			return
		}
	}

	// jika kategori tidak ditemukan
	errorResponse(w, "Category not found", http.StatusNotFound)
}

func main() {
	// health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		successResponse(w, "API is running", nil, http.StatusOK)
	})

	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getAllProducts(w, r)
		case "POST":
			createProduct(w, r)
		}
	})

	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getProductById(w, r)
		case "PUT":
			updateProduct(w, r)
		case "DELETE":
			deleteProduct(w, r)
		}
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getAllCategories(w, r)
		case "POST":
			createCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCategoryById(w, r)
		case "PUT":
			updateCategory(w, r)
		case "DELETE":
			deleteCategory(w, r)
		}
	})

	fmt.Println("Server is running on port 8889")

	err := http.ListenAndServe(":8889", nil)
	if err != nil {
		fmt.Println("Error running server:", err)
	}
}

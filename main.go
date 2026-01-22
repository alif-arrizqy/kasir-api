package main

import (
	"fmt"
	"kasir-api/utils"
	"kasir-api/services"
	"net/http"
)

func main() {
	// health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.SuccessResponse(w, "API is running", nil, http.StatusOK)
	})

	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			services.GetAllProducts(w, r)
		case "POST":
			services.CreateProduct(w, r)
		}
	})

	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			services.GetProductById(w, r)
		case "PUT":
			services.UpdateProduct(w, r)
		case "DELETE":
			services.DeleteProduct(w, r)
		}
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			services.GetAllCategories(w, r)
		case "POST":
			services.CreateCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			services.GetCategoryById(w, r)
		case "PUT":
			services.UpdateCategory(w, r)
		case "DELETE":
			services.DeleteCategory(w, r)
		}
	})

	fmt.Println("Server is running on port 8889")

	err := http.ListenAndServe(":8889", nil)
	if err != nil {
		fmt.Println("Error running server:", err)
	}
}

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
var produk = []Product{
	{ID: 1, Nama: "Indomie", Harga: 7000, Stok: 10},
	{ID: 2, Nama: "Vit 100ml", Harga: 1000, Stok: 20},
	{ID: 3, Nama: "Kecap", Harga: 3000, Stok: 20},
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
			produk = append(produk[:i], produk[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Product not found"})
}

func main() {
	// health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "message": "API is running"})
	})

	// GET localhost:8889/api/produk
	// POST localhost:8889/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		case "POST":
			// baca data dari request
			var produkBaru Product
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			// masukin data ke dalam array produk
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// GET localhost:8889/api/produk/{id}
	// PUT localhost:8889/api/produk/{id}
	// DELETE localhost:8889/api/produk/{id}
	http.HandleFunc("/api/produk/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getProductById(w, r)
		case "PUT":
			updateProduct(w, r)
		case "DELETE":
			deleteProduct(w, r)
		}
	})

	fmt.Println("Server is running on port 8889")

	err := http.ListenAndServe(":8889", nil)
	if err != nil {
		fmt.Println("Error running server:", err)
	}
}

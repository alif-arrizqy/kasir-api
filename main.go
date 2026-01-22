package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// blueprint untuk produk
type Produk struct {
	ID    int    `json:"id" validate:"required"`
	Nama  string `json:"nama" validate:"required"`
	Harga int    `json:"harga" validate:"required"`
	Stok  int    `json:"stok" validate:"required"`
}

// in-memory storage
var produk = []Produk{
	{ID: 1, Nama: "Indomie", Harga: 7000, Stok: 10},
	{ID: 2, Nama: "Vit 100ml", Harga: 1000, Stok: 20},
	{ID: 3, Nama: "Kecap", Harga: 3000, Stok: 20},
}

func getProdukById(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
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
	json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Produk not found"})
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// convert id to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	// loop produk, cari id yang sesuai, ganti data sesuai dengan data dari request
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}

	// jika produk tidak ditemukan
	json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Produk not found"})
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	
	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
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

	json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Produk not found"})
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
			var produkBaru Produk
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
			getProdukById(w, r)
		case "PUT":
			updateProduk(w, r)
		case "DELETE":
			deleteProduk(w, r)
		}
	})

	fmt.Println("Server is running on port 8889")

	err := http.ListenAndServe(":8889", nil)
	if err != nil {
		fmt.Println("Error running server:", err)
	}
}

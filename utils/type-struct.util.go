package utils

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

// in-memory storage
var Products = []Product{
	{ID: 1, Name: "Indomie", Price: 7000, Stock: 10},
	{ID: 2, Name: "Vit 100ml", Price: 1000, Stock: 20},
	{ID: 3, Name: "Kecap", Price: 3000, Stock: 20},
}

var Categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Makanan ringan"},
	{ID: 2, Name: "Minuman", Description: "Minuman segar"},
	{ID: 3, Name: "Lainnya", Description: "Lainnya"},
}

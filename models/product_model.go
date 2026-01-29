package models

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	CategoriesID int    `json:"categories_id"`
}

type ProductResponse struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Price    int       `json:"price"`
	Stock    int       `json:"stock"`
	Category *Category `json:"category"`
}

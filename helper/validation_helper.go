package helper

import "kasir-api/models"

// Validation Handler
func ValidateProduct(p models.Product) (string, bool) {
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
func ValidateCategory(c models.Category) (string, bool) {
	// Validasi Name
	if c.Name == "" {
		return "Field 'name' is required", false
	}

	// Semua validasi passed
	return "", true
}

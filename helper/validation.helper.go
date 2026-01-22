package helper

import "kasir-api/utils"

// Validation Handler
func ValidateProduct(p utils.Product) (string, bool) {
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
func ValidateCategory(c utils.Category) (string, bool) {
	// Validasi Name
	if c.Name == "" {
		return "Field 'name' is required", false
	}

	// Semua validasi passed
	return "", true
}

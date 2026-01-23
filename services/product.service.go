package services

import (
	"errors"
	"kasir-api/utils"
)

func GetAllProducts() []utils.Product {
	return utils.Products
}

func CreateProduct(newProduct utils.Product) utils.Product {
	// insert data ke dalam array produk
	newProduct.ID = len(utils.Products) + 1       // generate id auto increment
	utils.Products = append(utils.Products, newProduct) // insert data ke dalam array produk
	return newProduct
}

func GetProductById(id int) (utils.Product, error) {
	// cari produk dengan id yang sesuai
	for _, p := range utils.Products {
		if p.ID == id {
			return p, nil
		}
	}

	return utils.Product{}, errors.New("product not found")
}

func UpdateProduct(id int, updateProduct utils.Product) (utils.Product, error) {
	// loop produk, cari id yang sesuai, ganti data sesuai dengan data dari request
	for i := range utils.Products {
		if utils.Products[i].ID == id {
			updateProduct.ID = id
			utils.Products[i] = updateProduct

			return updateProduct, nil
		}
	}

	return utils.Product{}, errors.New("product not found")
}

func DeleteProduct(id int) error {
	// loop produk cari ID, dapet index yang mau dihapus
	for i, p := range utils.Products {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			utils.Products = append(utils.Products[:i], utils.Products[i+1:]...)
			return nil
		}
	}

	return errors.New("product not found")
}

package services

import (
	"errors"
	"kasir-api/utils"
)

func GetAllCategories() []utils.Category {
	return utils.Categories
}

func GetCategoryById(id int) (utils.Category, error) {
	for _, c := range utils.Categories {
		if c.ID == id {
			return c, nil
		}
	}
	return utils.Category{}, errors.New("category not found")
}

func CreateCategory(newCategory utils.Category) utils.Category {
	// masukin data ke dalam array kategori
	newCategory.ID = len(utils.Categories) + 1
	utils.Categories = append(utils.Categories, newCategory)
	return newCategory
}

func UpdateCategory(id int, updateCategory utils.Category) (utils.Category, error) {
	// loop categories, cari id yang sesuai, update data
	for i := range utils.Categories {
		if utils.Categories[i].ID == id {
			updateCategory.ID = id // pastikan ID tetap sama
			utils.Categories[i] = updateCategory

			return updateCategory, nil
		}
	}
	
	return utils.Category{}, errors.New("category not found")
}

func DeleteCategory(id int) error {
	// loop categories, cari ID, hapus dari slice
	for i, _ := range utils.Categories {
		if utils.Categories[i].ID == id {
			// hapus dengan append slice sebelum dan sesudah index
			utils.Categories = append(utils.Categories[:i], utils.Categories[i+1:]...)
			return nil
		}
	}
	
	return errors.New("category not found")
}

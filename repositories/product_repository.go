package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type ProductRepository struct {
	db           *sql.DB
	categoryRepo *CategoryRepository
}

func NewProductRepository(db *sql.DB, categoryRepo *CategoryRepository) *ProductRepository {
	return &ProductRepository{db: db, categoryRepo: categoryRepo}
}

func (repo *ProductRepository) GetAll() ([]models.ProductResponse, error) {
	// INNER JOIN: hanya product yang punya category valid (sesuai validasi Create/Update)
	query := `SELECT p.id, p.name, p.price, p.stock, c.id, c.name, c.description
	FROM products p
	INNER JOIN categories c ON p.categories_id = c.id
	ORDER BY p.id ASC`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productResponses := make([]models.ProductResponse, 0)
	for rows.Next() {
		var p models.Product
		var category models.Category
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		productResponses = append(productResponses, models.ProductResponse{
			ID:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			Stock:    p.Stock,
			Category: &category,
		})
	}
	return productResponses, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	// find category by CategoriesID
	category, err := repo.categoryRepo.GetByID(product.CategoriesID)
	if err != nil {
		return err
	}
	if category == nil {
		return errors.New("Category not found")
	}

	query := "INSERT INTO products (name, price, stock, categories_id) VALUES ($1, $2, $3, $4) RETURNING id"
	errInsert := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoriesID).Scan(&product.ID)
	return errInsert
}

// GetByID - ambil produk by ID
func (repo *ProductRepository) GetByID(id int) (*models.ProductResponse, error) {
	// INNER JOIN: hanya product yang punya category valid (sesuai validasi Create/Update)
	query := `SELECT p.id, p.name, p.price, p.stock, c.id, c.name, c.description
	FROM products p
	INNER JOIN categories c ON p.categories_id = c.id
	WHERE p.id = $1`

	var p models.Product
	var category models.Category
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &category.ID, &category.Name, &category.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("Product not found")
	}
	if err != nil {
		return nil, err
	}
	return &models.ProductResponse{
		ID:       p.ID,
		Name:     p.Name,
		Price:    p.Price,
		Stock:    p.Stock,
		Category: &category,
	}, nil
}

func (repo *ProductRepository) Update(product *models.Product) error {
	// find category by CategoriesID
	category, err := repo.categoryRepo.GetByID(product.CategoriesID)
	if err != nil {
		return err
	}
	if category == nil {
		return errors.New("Category not found")
	}

	query := "UPDATE products SET name = $1, price = $2, stock = $3, categories_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoriesID, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Product not found")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Product not found")
	}

	return err
}

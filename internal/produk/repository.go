package produk

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(name string) ([]ProductResponse, error) {
	query := "SELECT p.id, p.name as product_name, p.price, p.stock, jsonb_build_object('id', c.id,'name', c.name) AS category FROM products p JOIN categories c ON c.id = p.category_id"

	args := []interface{}{}
	if name != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+name+"%")
	}

	rows, err := repo.db.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Inisialisasi sebuah list dinamis
	products := make([]ProductResponse, 0)

	for rows.Next() {
		var p ProductResponse
		var categoryJSON []byte
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &categoryJSON); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(categoryJSON, &p.Category); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *ProductRepository) Create(p *Product) (string, error) {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"

	var newID string
	err := repo.db.QueryRow(query, p.Name, p.Price, p.Stock, p.CategoryID).Scan(&newID)
	if err != nil {
		return "", err
	}

	return newID, nil
}

func (repo *ProductRepository) GetByID(id string) (*ProductResponse, error) {
	log.Printf("ID : %s", id)
	query := "SELECT p.id, p.name as product_name, p.price, p.stock, jsonb_build_object('id', c.id,'name', c.name) AS category FROM products p JOIN categories c ON c.id = p.category_id WHERE p.id=$1"

	var p ProductResponse
	var categoryJSON []byte
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &categoryJSON)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Product doesn't exist")
		}

		return nil, err
	}
	// Parse JSONB -> struct
	if err := json.Unmarshal(categoryJSON, &p.Category); err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *ProductRepository) Update(id string, p *Product) error {
	query := "UPDATE products SET name=$1, price=$2, stock=$3, category_id=$4 WHERE id=$5"
	result, err := repo.db.Exec(query, p.Name, p.Price, p.Stock, p.CategoryID, id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Products doesn't exist")
	}

	return nil
}

func (repo *ProductRepository) Delete(id string) error {
	query := "DELETE FROM products WHERE id=$1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if rows == 0 {
		return errors.New("Products doesn't exist")
	}
	return err
}

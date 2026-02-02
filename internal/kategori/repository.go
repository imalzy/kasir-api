package kategori

import (
	"database/sql"
	"errors"
	"log"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAll() ([]Category, error) {
	query := "SELECT p.id, p.name, p.description FROM categories p"
	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Inisialisasi sebuah list dinamis
	categories := make([]Category, 0)

	for rows.Next() {
		var p Category
		if err := rows.Scan(&p.ID, &p.Name, &p.Description); err != nil {
			return nil, err
		}

		categories = append(categories, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (repo *CategoryRepository) Create(p *Category) (string, error) {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"

	var newID string
	err := repo.db.QueryRow(query, p.Name, p.Description).Scan(&newID)
	if err != nil {
		return "", err
	}

	return newID, nil
}

func (repo *CategoryRepository) GetByID(id string) (*Category, error) {
	log.Printf("ID : %s", id)
	query := "SELECT p.id, p.name, p.description categories p WHERE p.id=$1"

	var p Category
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Product doesn't exist")
		}

		return nil, err
	}

	return &p, nil
}

func (repo *CategoryRepository) Update(id string, p *Category) error {
	query := "UPDATE categories SET name=$1, description=$2, WHERE id=$3"
	result, err := repo.db.Exec(query, p.Name, p.Description, id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Category doesn't exist")
	}

	return nil
}

func (repo *CategoryRepository) Delete(id string) error {
	query := "DELETE FROM categories WHERE id=$1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if rows == 0 {
		return errors.New("Category doesn't exist")
	}
	return err
}

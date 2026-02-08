package transaction

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []CheckoutItem) (*Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	productIDs := make([]string, len(items))
	itemMap := make(map[string]CheckoutItem)
	for i, item := range items {
		productIDs[i] = item.ProductID
		itemMap[item.ProductID] = item
	}

	rows, err := tx.Query(
		"SELECT id, name, price, stock FROM products WHERE id = ANY($1::uuid[]) FOR UPDATE",
		pq.Array(productIDs),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type ProductData struct {
		ID    string
		Name  string
		Price int
		Stock int
	}
	productsMap := make(map[string]ProductData)

	for rows.Next() {
		var p ProductData
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
			return nil, err
		}
		productsMap[p.ID] = p
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	rows.Close()

	totalAmount := 0
	details := make([]TransactionDetail, 0)

	for _, item := range items {
		product, exists := productsMap[item.ProductID]
		if !exists {
			return nil, fmt.Errorf("produk dengan ID %s tidak ditemukan", item.ProductID)
		}

		if product.Stock < item.Quantity {
			return nil, fmt.Errorf("stok produk %s tidak mencukupi (sisa: %d)", product.Name, product.Stock)
		}

		subtotal := product.Price * item.Quantity
		totalAmount += subtotal

		result, err := tx.Exec(
			"UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1",
			item.Quantity, item.ProductID,
		)
		if err != nil {
			return nil, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return nil, err
		}
		if rowsAffected == 0 {
			return nil, fmt.Errorf("gagal update stok produk %s (kemungkinan concurrent modification)", product.Name)
		}

		details = append(details, TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: product.Name,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID string
	err = tx.QueryRow(
		"INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id",
		totalAmount,
	).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	if len(details) > 0 {
		query := "INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES "
		values := []interface{}{}
		for i, d := range details {
			p := i * 4
			query += fmt.Sprintf("($%d, $%d, $%d, $%d),", p+1, p+2, p+3, p+4)
			values = append(values, transactionID, d.ProductID, d.Quantity, d.Subtotal)
		}
		query = query[:len(query)-1]

		if _, err := tx.Exec(query, values...); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

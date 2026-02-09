package report

import (
	"database/sql"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetReportData(filter ReportFilter) (*ReportResponse, error) {
	// 1. Define the query with placeholders
	query := `
    WITH Totals AS (
        SELECT SUM(total_amount) as rev, COUNT(id) as trans
        FROM transactions WHERE created_at BETWEEN $1 AND $2
    ),
    ProductSales AS (
        SELECT p.name, SUM(td.quantity) as qty
        FROM transactions t
        JOIN transaction_details td ON t.id = td.transaction_id
        JOIN products p ON td.product_id = p.id
        WHERE t.created_at BETWEEN $1 AND $2
        GROUP BY p.name ORDER BY qty DESC LIMIT 1
    )
    SELECT rev, trans, name, qty FROM Totals, ProductSales`

	var resp ReportResponse

	// Go maps filter.StartDate to $1 and filter.EndDate to $2
	err := r.db.QueryRow(query, filter.StartDate, filter.EndDate).Scan(
		&resp.TotalRevenue,
		&resp.TotalTransaksi,
		&resp.ProdukTerlaris.Nama,
		&resp.ProdukTerlaris.QtyTerjual,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return an empty report instead of an error
			return &ReportResponse{
				TotalRevenue:   0,
				TotalTransaksi: 0,
			}, nil
		}
		return nil, err
	}

	return &resp, nil
}

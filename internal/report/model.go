package report

type ReportResponse struct {
	TotalRevenue   float64 `json:"total_revenue"`
	TotalTransaksi int     `json:"total_transaksi"`
	ProdukTerlaris struct {
		Nama       string `json:"nama"`
		QtyTerjual int    `json:"qty_terjual"`
	} `json:"produk_terlaris"`
}

type ReportFilter struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

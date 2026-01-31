package produk

type Product struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID string `json:"category_id"`
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductResponse struct {
	ID       string           `json:"id"`
	Name     string           `json:"product_name"`
	Price    int              `json:"price"`
	Stock    int              `json:"stock"`
	Category CategoryResponse `json:"category"`
}

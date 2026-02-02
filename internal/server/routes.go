package server

import "net/http"

func SetupRoutes(h Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Api is Running"))
	})

	if h.Product != nil {
		mux.HandleFunc("/api/product", h.Product.HandleProducts)
		mux.HandleFunc("/api/product/", h.Product.HandleProductByID)
	}

	if h.Category != nil {
		mux.HandleFunc("/api/category", h.Category.HandleCategories)
		mux.HandleFunc("/api/category/", h.Category.HandleCategoryByID)
	}

	return mux
}

package server

import (
	"encoding/json"
	"kasir-api/internal/kategori"
	"kasir-api/internal/produk"
	"log"
	"net/http"
)

func Start(addr string) {
	// Produk routes
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			produk.GetProdukByID(w, r)
		case http.MethodPut:
			produk.UpdateProduk(w, r)
		case http.MethodDelete:
			produk.DeleteProduk(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			produk.ListProduk(w, r)
		case http.MethodPost:
			produk.CreateProduk(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/kategori/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			kategori.GetKategoriByID(w, r)
		case http.MethodPut:
			kategori.UpdateKategori(w, r)
		case http.MethodDelete:
			kategori.DeleteKategori(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/kategori", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			kategori.ListKategori(w, r)
		case http.MethodPost:
			kategori.CreateKategori(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "API Running",
		}); err != nil {
			log.Printf("Failed to encode JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Printf("Server running on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

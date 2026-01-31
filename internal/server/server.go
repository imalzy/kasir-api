package server

import (
	"kasir-api/internal/produk"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	Product *produk.ProductHandler
	// Category *kategori.Category
	// Order    *transaksi.OrderHandler
}

func Start(addr string, h Handlers) {
	router := SetupRoutes(h)

	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Starting server at %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v", addr, err)
	}
}

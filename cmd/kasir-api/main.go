package main

import (
	"database/sql"
	"kasir-api/internal/produk"
	"kasir-api/internal/server"
	"kasir-api/utils"
	"log"
)

var version = "dev"

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db, err := utils.InitDb(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed connect into database")
	}

	defer db.Close()

	h := initHandlers(db)
	server.Start(":"+cfg.Port, h)
}

func initHandlers(db *sql.DB) server.Handlers {
	pRepo := produk.NewProductRepository(db)
	pService := produk.NewProductService(pRepo)
	pHandler := produk.NewProductHandler(pService)

	return server.Handlers{
		Product: pHandler,
	}
}

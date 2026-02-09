package main

import (
	"database/sql"
	"kasir-api/internal/kategori"
	"kasir-api/internal/produk"
	"kasir-api/internal/report"
	"kasir-api/internal/server"
	"kasir-api/internal/transaction"
	"kasir-api/utils"
	"log"
)

var version = "dev"

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db, err := utils.InitDb(cfg.DatabaseURL)
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

	cRepo := kategori.NewCategoryRepository(db)
	cService := kategori.NewCategoryService(cRepo)
	cHandler := kategori.NewProductHandler(cService)

	tRepo := transaction.NewTransactionRepository(db)
	tService := transaction.NewTransactionService(tRepo)
	tHandler := transaction.NewTransactionHandler(tService)

	rRepo := report.NewReportRepository(db)
	rService := report.NewReportService(rRepo)
	rHandler := report.NewReportHandler(rService)

	return server.Handlers{
		Product:     pHandler,
		Category:    cHandler,
		Transaction: tHandler,
		Report:      rHandler,
	}
}

package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xray-post-test/config"
)

func SambungKePostgres(cfg *config.Konfigurasi) *sql.DB {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatalf("Gagal membuka koneksi database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	fmt.Println("Berhasil terhubung ke PostgreSQL")
	return db
}

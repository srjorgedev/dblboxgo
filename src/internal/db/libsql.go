package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func OpenConnection() (*sql.DB, error) {
	url := os.Getenv("TURSO_URL")
	token := os.Getenv("TURSO_TOKEN")

	if url == "" || token == "" {
		return nil, fmt.Errorf("Missing TURSO_URL or TURSO_TOKEN, it may not set")
	}

	dsn := fmt.Sprintf("%s?authToken=%s", url, token)

	db, err := sql.Open("libsql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	// Ping to test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return db, nil
}

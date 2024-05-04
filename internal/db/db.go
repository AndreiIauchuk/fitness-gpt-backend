package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose/v3"
    _ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dialect = "pgx"
)

// Init initiates DB connection and applies the DB migration files if connection is established.
func Init() {
	var db *sql.DB
	defer db.Close()

	db = initConn()
	migrate(db)
}

func initConn() *sql.DB {
	connStr := fmt.Sprintf(
		"dbname=%s host=%s port=%s user=%s password=%s",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := goose.OpenDBWithDriver(dialect, connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database, because: %s", err)
	}

	log.Println("Connected to database and pinged it successfully!")
	return db
}

func migrate(db *sql.DB) {
	if err := goose.Up(db, "../migrations"); err != nil {
		log.Fatalf("Unable to apply DB migrations: %s", err)
	}

	log.Println("Migrated database successfully!")
}

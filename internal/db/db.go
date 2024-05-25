package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dialect = "pgx"
)

// Init initiates DB connection and applies the DB migration files if connection is established.
func Init() {
	db := initConn()
	defer db.Close()

	configGoose()
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

func configGoose() {
	goose.SetTableName(os.Getenv("DB_SCHEMA") + ".goose_db_version")
}

func migrate(db *sql.DB) {
	migrations_dir := fmt.Sprintf("%s/migrations", os.Getenv("WORK_DIR"))
	if err := goose.Up(db, migrations_dir); err != nil {
		log.Fatalf("Unable to apply DB migrations: %s", err)
	}

	log.Println("Migrated database successfully!")
}

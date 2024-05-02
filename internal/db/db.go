package db

import (
	"context"
	"database/sql"
	"embed"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Init() {
    var db *sql.DB;

    url := "dbname=postgres_db host=localhost port=5432 user=postgres password='postgres'"
    // TODO conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

    conn, err := pgx.Connect(context.Background(), url);
    log.Println("HI ALL")
	if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err);
	}
	defer conn.Close(context.Background())

    /* driver := "postgres" // Change this according to your database driver
    dbstring := "user=your_user dbname=your_db_name sslmode=disable" // Change this according to your database connection string 
    dir := "db/migrations"*/

    goose.SetBaseFS(embedMigrations);
    if err := goose.SetDialect("postgres"); err != nil {
        panic(err)
    }

    if err := goose.Up(db, "migrations"); err != nil {
        panic(err)
    }

    // Run the migrations
   /*  err := goose.Up(dbstring, dir)
    if err != nil {
        log.Fatalf("Error migrating the database: %v", err)
    }

    log.Println("Database migration successful!") */
}
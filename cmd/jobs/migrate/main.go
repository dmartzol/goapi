package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dbname := os.Getenv("PGDATABASE")
	dbusername := os.Getenv("PGUSER")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhostname := os.Getenv("PGHOST")

	dbURL := fmt.Sprintf("postgres://%s:5432/%s?user=%s&password=%s&sslmode=disable", dbhostname, dbname, dbusername, dbpassword)
	fmt.Println("Waiting for:", dbusername+"@tcp("+dbhostname+":)/"+dbname)

	var db *sql.DB
	for {
		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatalf("failed to open DB: %+v", err)
		}

		err = db.Ping()
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create driver: %+v", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %+v", err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("failed to close DB: %v\n", err)
		}
	}()
}

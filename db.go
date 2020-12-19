package bookreviews

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

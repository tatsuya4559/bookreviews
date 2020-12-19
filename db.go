package bookreviews

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func InitializeDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

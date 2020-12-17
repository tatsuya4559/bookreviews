package models

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func InitDB() error {
	var err error

	db, err = sqlx.Open("sqlite3", os.Getenv("DATABASE"))
	if err != nil {
		return err
	}

	return db.Ping()
}

type Book struct {
	ISBN      string
	Title     string
	Publisher string
}

func (b *Book) String() string {
	return fmt.Sprintf("%s, %s, %s", b.ISBN, b.Title, b.Publisher)
}

func AllBooks() ([]*Book, error) {
	bks := make([]*Book, 0)

	err := db.Select(&bks, `SELECT * FROM book`)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

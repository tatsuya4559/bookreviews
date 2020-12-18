package bookreviews

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/tatsuya4559/bookreviews/models"

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

type BookRepository interface {
	AllBooks() ([]*models.Book, error)
}

/***************************************************************/
type mockBookRepository struct {
	store map[uint]*models.Book
}

func newMockBookRepo() BookRepository {
	store := make(map[uint]*models.Book)
	return &mockBookRepository{store}
}

func (r *mockBookRepository) AllBooks() ([]*models.Book, error) {
	bks := make([]*models.Book, 0)
	for _, v := range r.store {
		bks = append(bks, v)
	}
	return bks, nil
}

/***************************************************************/
type bookRepository struct {
	db *sqlx.DB
}

func (r *bookRepository) AllBooks() ([]*models.Book, error) {
	bks := make([]*models.Book, 0)

	err := r.db.Select(&bks, `SELECT * FROM book`)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

func NewBookRepository(db *sqlx.DB) BookRepository {
	return &bookRepository{db: db}
}

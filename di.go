package bookreviews

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/dig"
)

var container *dig.Container

func BuildContainer() {
	container = dig.New()

	// *sqlx.DB initialized once
	db := InitializeDB()
	container.Provide(func() *sqlx.DB { return db })

	container.Provide(NewBookRepository)
}

package bookreviews

import (
	"go.uber.org/dig"
)

var container *dig.Container

func BuildContainer() {
	container = dig.New()

	container.Provide(InitializeDB)
	container.Provide(NewBookRepository)
}

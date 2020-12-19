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

func Invoke(function interface{}) {
	// 中身のクロージャが値を返すように作るつもりは無いけど
	// 型定義で縛れないので一応戻り値を見ておく
	err := container.Invoke(function)
	if err != nil {
		panic(err)
	}
}

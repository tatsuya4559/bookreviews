package models

import (
	"fmt"
)

type Book struct {
	ISBN      uint
	Title     string
	Publisher string
}

func (b *Book) String() string {
	return fmt.Sprintf("%d, %s, %s", b.ISBN, b.Title, b.Publisher)
}

package models

import (
	"fmt"
)

type Book struct {
	ISBN      string
	Title     string
	Publisher string
}

func (b *Book) String() string {
	return fmt.Sprintf("%s, %s, %s", b.ISBN, b.Title, b.Publisher)
}

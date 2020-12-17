package bookreviews

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tatsuya4559/bookreviews/models"
)

func booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s\n", bk.String())
	}
}

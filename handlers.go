package bookreviews

import (
	"fmt"
	"log"
	"net/http"
)

func booksIndex(w http.ResponseWriter, r *http.Request) {
	container.Invoke(func(repo BookRepository) {
		// ここでサービス層がDIされてきたりするけど、
		// それは別途テストされているからハンドラのテストではない
		// もう１階層あるとありがたみが分かる
		bks, err := repo.AllBooks()
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		for _, bk := range bks {
			fmt.Fprintf(w, "%s\n", bk.String())
		}
	})
}

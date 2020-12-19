package bookreviews

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFuncに値を渡すには
// context, グローバル変数, ハンドラーをメソッドにして構造体フィールドに持たせる
// が考えられる。
// contextはライフタイム管理のためのものなのでNG
// 構造体フィールドにするとその構造体が巨大になるし、どのハンドラが何に依存しているのかわかりにくい
// というわけでDIコンテナだけはグローバル変数に持たせて妥協してハンドラの中でInvokeを使ってみた
// コンテナだけならグローバルでもいいかな
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

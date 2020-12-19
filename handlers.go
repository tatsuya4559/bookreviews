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
// コンテナはパッケージをわけてグローバル変数とする
func booksIndex(w http.ResponseWriter, r *http.Request) {
	// What handlers do is
	//     parse request
	//     call usecases
	//     error handling
	//     write response
	Invoke(func(repo BookRepository) {
		// ここでサービス層がDIされてきたりするけど、
		// それは別途テストされているからハンドラのテストではない
		// もう１階層あるとありがたみが分かる
		bks, err := repo.All()
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

// What presentation layer do is
//     validation
//     classify erros - 400 or 500
//     call logic

// It's difficult to definitely separate presentation and usecase.
// Must separate if application has multiple gateway: REST, gRPC.
// But in most case, application has only one gateway.
// 面倒だから別の出力がある部分だけpresentationをかませたら良さげ

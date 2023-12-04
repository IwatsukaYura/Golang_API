## Golang API　チュートリアル
### 概要
GolangでAPIサーバー構築のチュートリアルです。
なにか詰まったところがあればここに戻ってくる。

### パスパラメータの取得
ルーティングに「gorilla/mux」を使う際は
```
r := mux.NewRouter()

r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodPost)
```
上記のように{}の中に変数を設定＋正規表現を用いると入力を制限できる。
これをハンドラーが受け取り
```
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid article id", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article %d\n", articleID)
	io.WriteString(w, resString)
}
```
mux.Vars(req)でURLに含まれるパスパラメータを全てmap型で返すため、その中の["id"]を取り出している。この返り値はstringであるから
strconv.Atoiでstring→intに変換をしており、変換ができないもの(整数以外)はエラー処理
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


### クエリパラメータの取得
クエリパラメータとは・・「http://localhost:8080/list?aa=11&&bb=12」などのようにURLに含まれる?以降の「aa=11」や「bb=12」などの部分のことであり、GETリクエストで取得するデータの範囲の指定などに利用することができる。

url.URL型にURLのさまざまなパラメータを取得できるメソッドは揃っている。(以下参照)

```
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	resString := fmt.Sprintf("Article list page %d\n", page)
	io.WriteString(w, resString)
}
```

req.URL.Query()でクエリパラメータ―をMap型で取り出す。その中に「page」というKeyがあればpageの一つ目のValueを採用し、返り値に渡している。
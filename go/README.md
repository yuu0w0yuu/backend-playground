# What's this?
Goバックエンド学習用のプレイグラウンド

# `net/http`を使ったAPIサーバの構築
## 構造
- ハンドラ（リクエストの処理関数）、ルーティング（URLパスとハンドラ関数のマッピング）、サーバ起動処理の三つ

## 実装パターン
### デフォルトマルチプレクサ
- シンプルで学習・検証用途向き
- `http.HandleFunc()`でルーティング（パスとハンドラのマッピング）をマルチプレクサに登録
    ```
    http.HandleFunc(<PATH>, <HANDLER_NAME>)
    ```
- `ListenAndServe`の第二引数にnilを渡すことでデフォルトマルチプレクサを利用する(https://pkg.go.dev/net/http#ListenAndServe)
    ```
    http.ListenAndServe(<ADDRESS>, <HANDLER>>)
    ```

### カスタムマルチプレクサ
- ルーティング設定が`mux`に閉じるため安全で再利用性が高い方式
- `http.NewServeMux()`で生成したマルチプレクサに`mux.HandleFunc`でルーティングを登録する。`ListenAndServe`にはハンドラとしてmuxを渡す
    ```
    mux := http.NewServeMux() 
    
    mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api/status", apiHandler)

    http.ListenAndServe(port, mux)
    ```


# Reference
- https://zenn.dev/hsaki/books/golang-httpserver-internal/viewer/intro
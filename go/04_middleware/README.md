# 内容
- gorillaを使用

- 起動
    ```
    go run main.go

    or

    air
    ```

# Logging
- リクエスト -> `http.ListenAndServe` -> Router(mux) -> useで登録したミドルウェア -> ハンドラ
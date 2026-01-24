# 内容
- airでLive Reload

- 起動
    ```
    go run main.go
    ```

- クエリパラメータの処理
    ```
    curl http://localhost:8080/api/greet?name=taro
    ```
- パスパラメータの処理
    ```
    curl http://localhost:8080/api/user/123
    ```
- JSONリクエストボディの処理
    ```
    url http://localhost:8080/api/echo -H 'Content-Type: application/json' -d "{"name": "taro", "mail": "example@example.com"}"
    ```
- エラーハンドリングの実装
- ミドルウェアでのロギング
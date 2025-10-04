# Dockerイメージ化
- ビルド・実行
```
docker build . -t go-hello
docker run -d -p 8080:8080 go-hello
```
- テスト
```
curl -i http://localhost:8080/api/hello
```


# Docker実行時の考慮事項
### glibc依存対応
- 1.20では仕様変わっている模様
https://go.dev/doc/go1.20#cgo
```
RUN CGO_ENABLED=0 GOOS=linux go build -o /goapp .
```

### PID考慮
- `SIGTERM`を適切に処理する必要があるアプリケーションの場合、Shell形式(シェルがPID1になる形式)ではなくExec形式で記述する
    ```
    # Shell形式(`sh`がPID1になる)
    ENTRYPOINT echo TEST

    # Exec形式（`echo`がPID1になる）
    ENTRYPOINT ["echo", "TEST"]
    ```
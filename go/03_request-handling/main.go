package main

import (
	"net/http"
	"fmt"

	"request-handling/handler"
)

func main() {
	//マルチプレクサを作成
	mux := http.NewServeMux()

	//ルーティング登録
	mux.HandleFunc("/api/greet", handler.GreetHandler) //クエリパラメータ処理
	mux.HandleFunc("/api/user/", handler.UserHandler) //パスパラメータ処理
	mux.HandleFunc("/api/echo", handler.EchoHandler) //リクエストボディ処理

	//サーバー起動
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
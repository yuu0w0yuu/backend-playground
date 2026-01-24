package main

import (
	"net/http"
	"fmt"

	"04_middleware/handler"
	"04_middleware/middleware"

	"github.com/gorilla/mux"
)

func main() {
	//ルータ作成
	fmt.Println("Starting server on :8080")
	router := mux.NewRouter()

	// ミドルウェア登録(http.Handlerを引数・返り値に持つ関数を登録)
	router.Use(middleware.LoggingMiddleware)

	//ルーティング登録
	router.HandleFunc("/api/greet", handler.GreetHandler) //クエリパラメータ処理
	router.HandleFunc("/api/user/{id}", handler.UserHandler) //パスパラメータ処理
	router.HandleFunc("/api/echo", handler.EchoHandler) //リクエストボディ処理

	//サーバー起動
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
package main

import (
	"fmt"
	"net/http"
)




func Handler(w http.ResponseWriter, r *http.Request) {
	//レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200 OK

	//レスポンスボディデータの生成
	response := `{"message": "Hello, World!"}`

	//レスポンスデータ書き込み
	w.Write([]byte(response))
}

func main() {
	//マルチプレクサを作成
	mux := http.NewServeMux()

	//ルーティング登録
	mux.HandleFunc("/api/hello", Handler)
	//ほかのルーティングもここで登録可能

	//サーバー起動
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// //////////
// // デフォルトマルチプレクサを使用したシンプルなHTTPサーバーのパターン
// //////////
// func main() {
// 	// ルーティング登録
// 	http.HandleFunc("/api/hello", Handler)
// 	//ほかのルーティングもここで登録可能

// 	// サーバー起動
// 	// 
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}
// }

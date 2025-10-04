package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

//レスポンス用の構造体定義
type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data	string `json:"data,omitempty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200 OK

	responseData := response{
		Status:  http.StatusOK,
		Message: "Hello, Docker!",
		Data:    "This is a sample response from a Go server running in a Docker container.",
	}

	//JSON形式でレスポンスデータを生成
	jsonBytes, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
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
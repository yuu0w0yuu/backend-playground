package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// クエリパラメータから名前を取得して挨拶を返すハンドラ
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	
	w.Write([]byte("hello," + name))
	}

// パスパラメータからユーザーIDを取得してユーザー情報を返すハンドラ
func UserHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/user/")

	userid, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user id: " + fmt.Sprintf("%d", userid)))
}

// リクエストボディからメッセージを取得してそのまま返すハンドラ
func EchoHandler(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
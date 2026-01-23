package main

import (
	"github.com/gorilla/mux"
	"go-middle/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Article Posted!")
}

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "List of Articles")
}

func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	responseString := fmt.Sprintf("Article No: %d", articleId)
	io.WriteString(w, responseString)
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Nice Posted!")
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Comment Posted!")
}

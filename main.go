package main

import (
	"log"
	"net/http"

	"github.com/IwatsukaYura/Golang_API/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/comments", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}

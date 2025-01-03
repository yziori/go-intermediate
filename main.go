package main

import (
	"log"
	"net/http"

	"github.com/yziori/go-intermediate/handlers"
)

func main() {
	// ハンドラとは
	// HTTPリクエストを受け取って、それに対するHTTPレスポンスの内容をコネクションに書き込む関数のこと

	// 定義したハンドラをサーバーで使用するように登録
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/article/comment", handlers.PostCommentHandler)

	// サーバー起動時のログを出力する
	log.Println("server start at port 8080")

	// サーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

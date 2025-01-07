package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 接続に使用するユーザー・パスワード・データベース名を定義する
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"

	// データベースに接続するためのアドレス文を定義
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// Open関数を用いてデータベースに接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}

	// プログラムが終了するときコネクションがcloseされるようにする
	defer db.Close()

	// Pingメソッドで疎通確認する
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect to DB")
	}
}

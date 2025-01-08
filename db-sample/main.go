package main

import (
	"database/sql"
	"dbsample/models"
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

	const sqlStr = `
		select title, contents, username, nice
		from articles;
	`
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName,
			&article.NickNum)

		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}

package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

// テスト全体で共有するsql.DBの変数を定義
// setup内で定義すると、テスト関数内で参照できないため、グローバル変数として定義
var testDB *sql.DB

// 全テスト共有の前処理を書く
func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// 全テスト共通の後処理を書く
func teardown() {
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}

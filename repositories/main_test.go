package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// テスト全体で共有するsql.DBの変数を定義
// setup内で定義すると、テスト関数内で参照できないため、グローバル変数として定義
var testDB *sql.DB

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// テストデータをセットアップする関数
func setupTestData() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb",
		"--password=docker", "-e", "source ./testdata/setupDB.sql")
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

// テストデータを削除する関数
func cleanupDB() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb",
		"--password=docker", "-e", "source ./testdata/cleanupDB.sql")
	err := cmd.Run()

	if err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	return nil
}

// 全テスト共有の前処理を書く
func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup")
		return err
	}
	return nil
}

// 全テスト共通の後処理を書く
func teardown() {
	cleanupDB()
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

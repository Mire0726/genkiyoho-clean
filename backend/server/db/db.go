package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"net/url"

	// "github.com/joho/godotenv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// const driverName = "mysql"

var Conn *sql.DB

// func init() {

// 	err := godotenv.Load() 
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	user := os.Getenv("MYSQL_USER")
// 	password := os.Getenv("MYSQL_PASSWORD")
// 	host := os.Getenv("MYSQL_HOST")
// 	port := os.Getenv("MYSQL_PORT")
// 	database := os.Getenv("MYSQL_DATABASE")
// 	charset := os.Getenv("MYSQL_CHARSET")
// 	parseTime := os.Getenv("MYSQL_PARSE_TIME")
// 	loc := os.Getenv("MYSQL_LOC")
	
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
// 		user, password, host, port, database, charset, parseTime, loc)
	
// 	Conn, err := sql.Open(driverName, dsn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := Conn.Ping(); err != nil {
// 		log.Fatal("Unable to connect to the database:", err)
// 	}
// }

func init() {
    jawsdbURL := os.Getenv("JAWSDB_URL")
    if jawsdbURL == "" {
        log.Fatal("JAWSDB_URL environment variable is not set")
    }

    // JawsDB の接続 URL を Go の標準形式に変換
    parsedURL, err := url.Parse(jawsdbURL)
    if err != nil {
        log.Fatal(err)
    }

    // ユーザ名とパスワードを抽出
    user := parsedURL.User.String()

    // ホスト名とポート番号を抽出
    host := parsedURL.Host

    // データベース名を抽出
    dbName := strings.TrimPrefix(parsedURL.Path, "/")

    // データソース名 (DSN) を構築
    dsn := fmt.Sprintf("%s@tcp(%s)/%s?parseTime=true", user, host, dbName)

    Conn, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    if err := Conn.Ping(); err != nil {
        log.Fatal("Unable to connect to the database:", err)
    }
}
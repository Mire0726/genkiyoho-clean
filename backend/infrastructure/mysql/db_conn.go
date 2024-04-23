package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

const driverName = "mysql"

var Conn *sql.DB

func ConnectToDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err) // エラーメッセージに詳細を追加
	}

	log.Println("loaded .env file")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")
	charset := os.Getenv("MYSQL_CHARSET")
	parseTime := os.Getenv("MYSQL_PARSE_TIME")
	loc := os.Getenv("MYSQL_LOC")
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		user, password, host, port, database, charset, parseTime, loc)
	
	Conn, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Fatal("cannnot sql.Open",err)
	}
	if err := Conn.Ping(); err != nil {
		log.Fatal("Unable to connect to the database:", err)
	}
	log.Println("Database connection established") // 成功ログ
	
	return Conn, nil
	
}
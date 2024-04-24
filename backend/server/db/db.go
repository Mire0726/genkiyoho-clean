package db

import (
	"database/sql"
	// "fmt"
	// "log"
	// "os"


	// "github.com/joho/godotenv"
	

	_ "github.com/go-sql-driver/mysql"
)

const driverName = "mysql"

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
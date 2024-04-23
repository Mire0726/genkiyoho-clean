package main

import (
	"flag"

	"fmt"
	"log"
	"os"

	db "github.com/Mire0726/Genkiyoho/backend/infrastructure/mysql"

	presentation "github.com/Mire0726/Genkiyoho/backend/handler"
)

func main() {
	db, err := db.ConnectToDB()
	if err != nil {
		log.Fatal("Could not initialize database:", err) // エラーハンドリング
	}

	// データベース接続が確立されていることを確認
	if db == nil {
		log.Fatal("Database connection is nil in main") // エラーログ
	}

	// 環境変数PORTからポート番号を取得。指定されていない場合はデフォルトで"8080"を使用。
	var defaultPort = "8080"
	var port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		flag.StringVar(&port, "addr", defaultPort, "default server port")
	}
	flag.Parse()

	// サーバーの設定と起動
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Listening on %s...\n", addr)
	presentation.StartServer(addr)
}

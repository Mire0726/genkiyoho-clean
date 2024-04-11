package main

import (
	"flag"

	"github.com/Mire0726/Genkiyoho/backend/server"
    "os"
    "fmt"
    "log"
	
)

func main() {
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
	server.Serve(addr)
}

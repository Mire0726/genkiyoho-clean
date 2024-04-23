package routes

import (
	"log"
	"net/http"
	"time"

	handler "github.com/Mire0726/Genkiyoho/backend/handler/user"
	"github.com/labstack/echo/v4"

	userDB "github.com/Mire0726/Genkiyoho/backend/infrastructure/mysql"
	userI "github.com/Mire0726/Genkiyoho/backend/infrastructure/repositories"

	_ "github.com/go-sql-driver/mysql"

	echomiddleware "github.com/labstack/echo/v4/middleware"
)

// サーバーの設定とルーティング
func StartServer(addr string) {
	e := echo.New()
    // データベース接続の初期化
	db, err := userDB.ConnectToDB() // 接続オブジェクトの生成
	if err != nil {
		log.Fatal("Database connection failed to initialize:", err) // エラーハンドリング
	}
	// 接続オブジェクトが正しく初期化されているか確認
	if db == nil {
		log.Fatal("Database connection is not initialized") // 二重チェック
	}
    // リポジトリとユースケースの初期化
	repo := userI.NewMySQLUserRepository(db) // リポジトリの生成
	if repo == nil {
		log.Fatal("Failed to initialize repository") // エラーハンドリング
	}
	// ルーティングの設定
	timeout := time.Second * 10
	userCreateHandler := handler.NewUserCreateHandler(repo, timeout) // ユーザー作成ハンドラーの生成

	// ミドルウェアの設定
	e.Use(echomiddleware.Recover()) // パニックリカバリー
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"}, // CORS設定
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
	
	// ルーティングの設定
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Genkiyoho!")
	})

	e.POST("/users/me", userCreateHandler.Create()) // ユーザー作成エンドポイント

	// サーバーの起動
	log.Printf("Server running on %s", addr) // サーバー起動のログ
	if err := e.Start(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err) // サーバー起動エラーのハンドリング
	}
}
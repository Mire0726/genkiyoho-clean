package server

import (
	"log"
	"net/http"

	"github.com/Mire0726/Genkiyoho/backend/server/handler"
	"github.com/Mire0726/Genkiyoho/backend/server/http/middleware"
	// "github.com/Mire0726/Genkiyoho/backend/server/weather"

	_ "github.com/go-sql-driver/mysql" // MySQLドライバーをインポート
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

// Serve はHTTPサーバを起動します。データベース接続を引数に追加。
func Serve(addr string) {
    e := echo.New()
    
// ミドルウェアの設定
    // panicが発生した場合の処理
	e.Use(echomiddleware.Recover())
	// CORSの設定
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
        Skipper:      echomiddleware.DefaultCORSConfig.Skipper,
        AllowOrigins: echomiddleware.DefaultCORSConfig.AllowOrigins,
        AllowMethods: echomiddleware.DefaultCORSConfig.AllowMethods,
        AllowHeaders: []string{"Content-Type", "Accept", "Origin", "X-Token", "Authorization"},
    }))
    

    // ルーティングの設定
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Welcome to Genkiyoho!")
    })
    e.POST("/users/me", handler.HandleUserCreate()) // ユーザ登録API
    e.POST("/users/login", handler.HandleUserLogin) // ログインAPI
    e.GET("/users", handler.HandleGetUser()) // ユーザ一覧取得API
    e.GET("conditions",handler.HandleConditionsGet()) // 条件一覧取得API
    e.GET("conditions/cycle",handler.HandleCycleConditionGet()) // サイクル条件一覧取得API
    e.GET("conditions/environment",handler.HandleEnvironmentConditionGet()) // 環境条件一覧取得API
    
    authAPI := e.Group("", middleware.AuthenticateMiddleware())
    authAPI.GET("/users/me", handler.HandleUserGet()) // ユーザ情報取得API
    authAPI.PUT("/users/me", handler.HandleUserUpdate())  // ユーザ情報更新API
    authAPI.POST("users/me/condition/cycle",handler.HandleCycleConditionCreate()) // サイクル条件登録API
    authAPI.POST("users/me/condition/environment",handler.HandleEnvironmentConditionCreate()) // 環境条件登録API
    authAPI.GET("users/me/condition",handler.HandleUserConditionGet()) // 
    authAPI.GET("users/me/condition/today/cycle",handler.HandleUserTodayCycleConditionGet) // 本日のサイクル条件取得API
    authAPI.GET("users/me/condition/today/environment",handler.HandleUserTodayEnvironmentConditionGet) // 本日の環境条件取得API
    authAPI.GET("users/me/condition/today",handler.HandleUserTodayConditionGet) // 本日の状態取得API
    authAPI.GET("users/me/condition/today/point",handler.HandleUserTodayPointGet) // 本日のポイント取得API
    /* ===== サーバの起動 ===== */

    log.Printf("Server running on %s", addr)
    if err := e.Start(addr); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}


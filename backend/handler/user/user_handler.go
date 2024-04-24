package handler

import (
	"net/http"

	userA "github.com/Mire0726/Genkiyoho/backend/application/user"
	userR "github.com/Mire0726/Genkiyoho/backend/domain/repositories"

	// userI "github.com/Mire0726/Genkiyoho/backend/infrastructure/repositories"
	"time"

	"github.com/labstack/echo/v4"
)


type UserCreateHandler struct {
    userUsecase userA.UserUseCase // ユースケースを保持
}
type UserLoginHandler struct {
    userUsecase userA.UserUseCase // ユースケースを保持
}

// UserCreateHandler のコンストラクタ
func NewUserCreateHandler(repo userR.UserRepository, timeout time.Duration) UserCreateHandler {
    userUseCase := userA.NewUserInteractor(repo, timeout) // ユースケースインスタンスを生成
    return UserCreateHandler{userUsecase: userUseCase} // インスタンスを返す
}

// ユーザー作成ハンドラー
func (handler *UserCreateHandler) Create() echo.HandlerFunc {
    userUseCase := handler.userUsecase // Declare the userUseCase variable
    return func(c echo.Context) error {
        var userInput userA.CreateUserInput // ユーザーモデル
        if err := c.Bind(&userInput); err != nil {
            return c.JSON(http.StatusBadRequest, "Invalid input data") // バインディングエラーの処理
        }

        // ビジネスロジック層にユーザー作成を依頼
        _, err := userUseCase.Execute(c.Request().Context(), userInput) // ユースケースのメソッドを呼び出す
        if err != nil {
            return c.JSON(http.StatusInternalServerError, "Error creating user") // エラーの処理
        }

        return c.JSON(http.StatusOK, "User created successfully") // 成功メッセージ
    }
}

// ユーザーログインハンドラー
func NewUserLoginHandler(repo userR.UserRepository, timeout time.Duration) UserLoginHandler {
    userUseCase := userA.NewUserInteractor(repo, timeout) // ユースケースインスタンスを生成
    return UserLoginHandler{userUsecase: userUseCase}
}

func (handler *UserLoginHandler) Login() echo.HandlerFunc {
    userUseCase := handler.userUsecase // Declare the userUseCase variable
    return func(c echo.Context) error {
        var userInput userA.LoginInput // ユーザーモデル
        if err := c.Bind(&userInput); err != nil {
            return c.JSON(http.StatusBadRequest, "Invalid input data") // バインディングエラーの処理
        }

        // ビジネスロジック層にユーザー作成を依頼
        _, err := userUseCase.Login(c.Request().Context(), userInput) // ユースケースのメソッドを呼び出す
        if err != nil {
            return c.JSON(http.StatusInternalServerError, "Error login user") // エラーの処理
        }
        return c.JSON(http.StatusOK, "User created successfully") // 成功メッセージ
    }
}
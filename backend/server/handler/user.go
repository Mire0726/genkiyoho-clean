package handler

import (
	"net/http"
	"github.com/google/uuid"
	"time"
	"net/mail"

	"github.com/Mire0726/Genkiyoho/backend/server/model"
	"github.com/labstack/echo/v4"
	"log"
	"github.com/Mire0726/Genkiyoho/backend/server/context/auth"
	"errors"
    "github.com/Mire0726/Genkiyoho/backend/server/db"
)

type userCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// validateEmail はメールアドレスが有効な形式かどうかを検証します。
func validateEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

// HandleUserCreate ユーザ登録処理のハンドラ
func HandleUserCreate() echo.HandlerFunc {
    return func(c echo.Context) error {
		req := &userCreateRequest{}
		if err := c.Bind(req); err != nil {
			return err
		}

    // メールアドレスの形式を検証
    if !validateEmail(req.Email) {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid email format")
    }

	// UUIDでユーザIDを生成する
	userID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

    // UUIDで認証トークンを生成
    authToken, err := uuid.NewRandom()
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate authentication token")
    }

    now := time.Now() // 現在の時刻
	// データベースにユーザデータを登録する
	if err := model.InsertUser(&model.User{
		ID:		userID.String(),
		AuthToken: authToken.String(),
        Email:     req.Email,
        Password:  req.Password,
        Name:      req.Name,
        CreatedAt: now,
        UpdatedAt: now,
	}); err != nil {
		return err
	}
	// 生成した認証トークンを返却
	return c.JSON(http.StatusOK, &userCreateResponse{Token: authToken.String()})
}
}

type userCreateResponse struct {
	Token string `json:"token"`
}

// GetUser 全ユーザ情報を取得するエンドポイントのハンドラ
func HandleGetUser() echo.HandlerFunc{
	return func(c echo.Context) error {

		users, err :=model.GetAllUsers()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// echo.ContextのJSONメソッドを使用してユーザリストをJSON形式で返す
		return c.JSON(http.StatusOK, users)
	}
}

// HandleUserGet 特定のユーザ情報を取得するエンドポイントのハンドラ
func HandleUserGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        // Contextから認証済みのユーザIDを取得
        ctx := c.Request().Context()
        userID := auth.GetUserIDFromContext(ctx)
        if userID == "" {
            return errors.New("userID is empty")
        }

        // ユーザデータの取得処理を実装
        user, err := model.SelectUserByPrimaryKey(userID)
        if err != nil {
            return err
        }
        if user == nil {
            return errors.New("user not found")
        }

        // レスポンスに必要な情報を詰めて返却
        return c.JSON(http.StatusOK, user)
    }
}

type userUpdateRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
// HandleUserUpdate ユーザ情報更新処理
func HandleUserUpdate() echo.HandlerFunc {
    return func(c echo.Context) error {

        // リクエストBodyから更新後情報を取得
        req := &userUpdateRequest{}
        if err := c.Bind(req); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body format")
        }
     
        // Contextから認証済みのユーザIDを取得
        ctx := c.Request().Context()
        userID := auth.GetUserIDFromContext(ctx)
        if userID == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
        }
        
        // 更新対象のユーザデータを取得（存在チェック）
        userData, err := model.SelectUserByPrimaryKey(userID)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user data")
        }
        if userData == nil {
            return echo.NewHTTPError(http.StatusNotFound, "User not found")
        }
        
        log.Println("line150")
        
        // リクエストから受け取ったデータでユーザ情報を更新
        userData.Name = req.Name
        userData.Email = req.Email
        userData.Password = req.Password // 実運用ではパスワードをハッシュ化

        // ユーザデータの更新処理
        if err := model.UpdateUserByPrimaryKey(userData); err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user data")
        }
        
        return c.NoContent(http.StatusOK)
    }
}

// HandleUserLogin はユーザーのログイン処理を行う
func HandleUserLogin(c echo.Context) error {
    var req loginRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
    }


    user, err := model.AuthenticateUser(db.Conn, req.Email, req.Password)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Authentication failed")
    }
    if user == nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
    }

    // 認証成功。トークン生成やセッション管理などの処理をここに追加

    return c.JSON(http.StatusOK, user)
}

type loginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

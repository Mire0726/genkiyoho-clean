package middleware

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/Mire0726/Genkiyoho/backend/server/context/auth"
	"github.com/Mire0726/Genkiyoho/backend/server/model"
)

// AuthenticateMiddleware ユーザ認証を行ってContextへユーザID情報を保存する
func AuthenticateMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			// リクエストヘッダからx-token(認証トークン)を取得
			token := c.Request().Header.Get("x-token")
			if token == "" {
				return errors.New("x-token is empty")
			}
			
			// データベースから認証トークンに紐づくユーザの情報を取得
			user, err := model.SelectUserByAuthToken(token)
			if err != nil {
				return err
			}
			if user == nil {
				return fmt.Errorf("user not found. token=%s", token)
			}
			// ユーザIDをContextへ保存して以降の処理に利用する
			c.SetRequest(c.Request().WithContext(auth.SetUserID(ctx, user.ID)))
			
			// 次の処理
			return next(c)
		}
	}
}

package user

import (
	"fmt"
	"time"
	"context"
	// "unicode/utf8"
	// 	errDomain "github.com/Mire0726/Genkiyoho/backend/domain/error"
)

type User struct {
    ID        string
    AuthToken string
    Email     string
    Password  string
    Name      string
    CreatedAt time.Time
    UpdatedAt time.Time
}


func NewUser(id, authToken, email, password, name string, createdAt, updatedAt time.Time) (*User, error) {
    if email == "" || password == "" {
        return nil, fmt.Errorf("email and password are required")
    }
	
    return &User{id, authToken, email, password, name, createdAt, updatedAt}, nil
}

//ユーザー情報に関する操作を抽象化したインターフェース
type UserRepository interface {
    Save(ctx context.Context, user *User) error
	//Saveメソッドは、context.Context型のctxと*User型のuserを引数に取り、error型の値を返します。

}


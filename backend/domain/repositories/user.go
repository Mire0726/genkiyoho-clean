package user

import (
    // "database/sql"
	"context"
    userDomain "github.com/Mire0726/Genkiyoho/backend/domain/model"
)
// UserRepository はユーザーの永続化を抽象化するインターフェイスです。
type UserRepository interface {
    Insert(ctx context.Context, user *userDomain.User) error
    // Save(ctx context.Context, user *userDomain.User) error
}

// UserService はユーザー関連のドメインサービスを提供します。
type UserService struct {
    userRepo UserRepository
}

// NewUserService は新しいUserServiceを生成します。
func NewUserService(repo UserRepository) *UserService {
    return &UserService{
        userRepo: repo,
    }
}


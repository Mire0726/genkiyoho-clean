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
}// UserRepository is an interface that abstracts the saving of User entities

// type MySQLUserRepository struct {
// 	db *sql.DB
// }

// func NewMySQLUserRepository(context.Context,db *sql.DB) UserRepository {
// 	return &MySQLUserRepository{db: db}
// }

// func (r *MySQLUserRepository) Save(user *userDomain.User) error {
// 	query := `
// 		INSERT INTO users (email, password, created_at, updated_at)
// 		VALUES (?, ?, NOW(), NOW())
// 	`
// 	_, err := r.db.Exec(query, user.Email, user.PasswordHash)
// 	if err != nil {
// 		return fmt.Errorf("failed to save user: %w", err)
// 	}
// 	return nil
// }

// func (r *MySQLUserRepository) FindByEmail(email string) (*entities.User, error) {
// 	// FindByEmailの実装は省略
// 	return nil, nil
// }

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	userDomain "github.com/Mire0726/Genkiyoho/backend/domain/model"
)

type MySQLUserRepository struct {
    DB *sql.DB
}


func (r *MySQLUserRepository) Insert(ctx context.Context, user *userDomain.User) error {
    log.Println("Come to Inserting user:", user)
    log.Println("DB:", r.DB)
    if r.DB == nil {
        return fmt.Errorf("database connection is nil")
    }
    log.Println("Inserting user:", user)


    if user == nil {
        return fmt.Errorf("user is nil")
    }
    log.Println("Executing SQL insert for user:", user) 
    _, err := r.DB.ExecContext(
        ctx,
        "INSERT INTO users (id, name, email, password, auth_token) VALUES (?, ?, ?, ?, ?)",
        user.ID, user.Name, user.Email, user.Password, user.AuthToken,)
    if err != nil {
        return fmt.Errorf("failed to insert user: %w", err)
    }
    log.Println("User inserted:", user)

    return nil
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	if db == nil {
		log.Fatal("Provided database connection is nil") // エラーハンドリング
	}
	return &MySQLUserRepository{DB: db} // 正しく初期化
}


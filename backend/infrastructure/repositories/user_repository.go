package repository

import (
	"context"
	"database/sql"

	userDomain "github.com/Mire0726/Genkiyoho/backend/domain/model"
	userR "github.com/Mire0726/Genkiyoho/backend/domain/repositories"
)

type MySQLUserRepository struct {
    DB *sql.DB
}


func (r *MySQLUserRepository) Insert(ctx context.Context, user *userDomain.User) error {
    _, err := r.DB.ExecContext(ctx, "INSERT INTO users (id, name, email, password, auth_token, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
        user.ID, user.Name, user.Email, user.Password, user.AuthToken, user.CreatedAt, user.UpdatedAt)
    if err != nil {
        return err
    }
    return nil
}

func NewMySQLUserRepository(db *sql.DB) userR.UserRepository {
    return &MySQLUserRepository{DB: db}
}


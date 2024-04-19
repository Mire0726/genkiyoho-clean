package application

import (
	"context"

	"time"

	userDomain "github.com/Mire0726/Genkiyoho/backend/domain/model"
	userR "github.com/Mire0726/Genkiyoho/backend/domain/repositories"
	"github.com/google/uuid"
)

type CreateUserUseCase interface {
	Execute(context.Context, CreateUserInput) (string, error) // Returns auth token
}

type CreateUserInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type createUserInteractor struct {
	repo       userR.UserRepository
	ctxTimeout time.Duration
}


func NewCreateUserInteractor(repo userR.UserRepository, timeout time.Duration) CreateUserUseCase {
	return &createUserInteractor{
		repo:       repo,
		ctxTimeout: timeout,
	}
}

func (interactor *createUserInteractor) Execute(ctx context.Context, input CreateUserInput) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, interactor.ctxTimeout)
	defer cancel()

	id := uuid.New().String()
	authToken := uuid.New().String()
	now := time.Now()

	user := userDomain.User{
		ID:        id,
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password, // Hash this password in real implementation
		AuthToken: authToken,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := interactor.repo.Insert(ctx, &user); err != nil {
		return "", err
	}

	return authToken, nil
}
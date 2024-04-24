package application

import (
	"context"
	"fmt"
	"time"
	"log"

	userDomain "github.com/Mire0726/Genkiyoho/backend/domain/model"
	userR "github.com/Mire0726/Genkiyoho/backend/domain/repositories"
	// userA "github.com/Mire0726/Genkiyoho/backend/infrastructure/repositories"

	"github.com/google/uuid"
)

type UserUseCase interface {
	Execute(context.Context, CreateUserInput) (string, error) // Returns auth token
	Login(context.Context, LoginInput) (string, error)
}

type CreateUserInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserInteractor struct {
	repo       userR.UserRepository
	ctxTimeout time.Duration
}
func NewUserInteractor(repo userR.UserRepository, timeout time.Duration) *UserInteractor {
	return &UserInteractor{
		repo:       repo,
		ctxTimeout: timeout,
	}
}

func (interactor *UserInteractor) Execute(ctx context.Context, input CreateUserInput) (string, error) {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return "", fmt.Errorf("invalid input data")
	}
	log.Println("Creating user with name:", input.Name)
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
	log.Println("User created:", user)

	if err := interactor.repo.Insert(ctx, &user); err != nil {
		return "error !", fmt.Errorf("error during user creation insert: %v", err)
	}
	log.Println("User inserted into database")

	return authToken, nil
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (interactor *UserInteractor) Login(ctx context.Context, input LoginInput) (string, error) {
	if input.Email == "" || input.Password == "" {
		return "", fmt.Errorf("invalid input data")
	}

	ctx, cancel := context.WithTimeout(ctx, interactor.ctxTimeout)
	defer cancel()

	user := userDomain.User{
		Email:    input.Email,
		Password: input.Password,
	}
	log.Println("User logged in:", user)

	if err := interactor.repo.Login(ctx, &user); err != nil {
		return "error !", fmt.Errorf("error during user login GET: %v", err)
	}
	log.Println("User logged in")

	return user.AuthToken, nil
}
package repository

import (
	"github.com/Mire0726/Genkiyoho/backend/domain/user"
)

type userRepository struct {}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Save(user *user.User) error {
	
// application/userusecase.go
package application

import (
	userDomain "github.com/Mire0726/Genkiyoho/backend/domain/user"
    // "time"
	// "github.com/google/uuid"
)

type SaveUserUseCase struct {
    UserRepo userDomain.UserRepository
}

//ユーザー情報を保存するためのユースケースを作成し、そのユースケースがユーザー情報の操作を行うためのリポジトリを持つようにする
func NewSaveUserUsecase(
    UserRepo userDomain.UserRepository,
) *SaveUserUseCase{
    return &SaveUserUseCase{
        UserRepo: UserRepo,
    }
}

//ユーザー情報を保存するためのユースケース
func (u *SaveUserUseCase)Run (
    id, authToken, email, password, name string,
) error {
    user, err := userDomain.NewUser(id, authToken, email, password, name)
    if err != nil {
        return err
    }
    return u.UserRepo.Save(user)
}

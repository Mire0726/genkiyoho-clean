// application/userusecase.go
package application

import (
	// userDomain "github.com/Mire0726/Genkiyoho/backend/domain/model"
    // "time"
    // "context"
	// "github.com/google/uuid"
)

// type SaveUserUseCase struct {
//     UserRepo userDomain.UserRepository
// }

// //ユーザー情報を保存するためのユースケースを作成し、そのユースケースがユーザー情報の操作を行うためのリポジトリを持つようにする
// func NewSaveUserUsecase(
//     UserRepo userDomain.UserRepository,
// ) *SaveUserUseCase{
//     return &SaveUserUseCase{
//         UserRepo: UserRepo,
//     }
// }
// //ユーザー情報を保存するためのユースケース
// //Runメソッドは"u *SaveUserUseCase"で定義されていて、id, authToken, email, password, nameを引数に取り、error型の値を返します。
// func (u *SaveUserUseCase)Run ( id, authToken, email, password, name string,) error {
//     user, err := userDomain.NewUser(id, authToken, email, password, name, time.Now(), time.Now())
//     if err != nil {
//         return err
//     }
//     ctx := context.Background()
//     return u.UserRepo.Save(ctx,user)
// }

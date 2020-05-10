/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package usecases

import (
	"github.com/kazekim/UserData/app/entities"
	"github.com/kazekim/UserData/app/repositories"
)

type UserUseCase interface {
	GetUser(id int64) (*entities.User, error)
	ListUser() (*[]entities.User, error)

	CreateUser(email, fullName, telephone string) error
	UpdateUser(id int64, email, fullName, telephone string) error
	DeleteUser(id int64) error
}

type defaultUserUseCase struct {
	userDBRepo repositories.UserDBRepository
}

func NewUserUseCase(userDBRepo repositories.UserDBRepository) UserUseCase {
	return &defaultUserUseCase{
		userDBRepo: userDBRepo,
	}
}

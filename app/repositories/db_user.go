/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package repositories

import (
	"github.com/kazekim/UserData/app/entities"
	"github.com/kazekim/UserData/internal/userdb/v1"
)

type UserDBRepository interface {
	GetUser(id int64) (*entities.User, error)
	ListUser() (*[]entities.User, error)
	CreateUser(email, fullName, telephone string) error
	UpdateUser(id int64, email, fullName, telephone string) error
	DeleteUser(id int64) error
}

type defaultUserDBRepository struct {
	client userdb.Client
}

func NewUserDBRepository(client userdb.Client) UserDBRepository {
	return &defaultUserDBRepository{
		client: client,
	}
}

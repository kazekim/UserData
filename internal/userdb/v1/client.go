/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package userdb

import (
	"database/sql"
	userdbentities "github.com/kazekim/UserData/internal/userdb/v1/entities"
	_ "github.com/mattn/go-sqlite3"
)

type Client interface {
	GetUser(id int64) (*userdbentities.User, error)
	List() (*[]userdbentities.User, error)
	CreateUser(user userdbentities.User) error
	CreateUserByField(email, fullName, telephone string) error
	UpdateUser(user *userdbentities.User) error
	DeleteUser(id int64) error
}

type defaultClient struct {
	db *sql.DB
}

func NewClient(db *sql.DB) Client {
	return &defaultClient{
		db: db,
	}
}

/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package repositories

import (
	"github.com/kazekim/UserData/app/entities"
	"github.com/mitchellh/mapstructure"
)

func (repo *defaultUserDBRepository) GetUser(id int64) (*entities.User, error) {

	user, err := repo.client.GetUser(id)
	if err != nil {
		return nil, err
	}
	var u entities.User
	err = mapstructure.Decode(user, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package repositories

import (
	"github.com/kazekim/UserData/app/entities"
	"github.com/mitchellh/mapstructure"
)

func (repo *defaultUserDBRepository) ListUser() (*[]entities.User, error) {

	users, err := repo.client.List()
	if err != nil {
		return nil, err
	}

	var us []entities.User
	for _, user := range *users {
		var u entities.User
		err = mapstructure.Decode(user, &u)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}

	return &us, nil
}

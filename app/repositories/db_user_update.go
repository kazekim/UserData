/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package repositories

import userdbentities "github.com/kazekim/UserData/internal/userdb/v1/entities"

func (repo *defaultUserDBRepository) UpdateUser(id int64, email, fullName, telephone string) error {
	user := userdbentities.User{
		ID:        id,
		Email:     email,
		Fullname:  fullName,
		Telephone: telephone,
	}
	return repo.client.UpdateUser(&user)
}

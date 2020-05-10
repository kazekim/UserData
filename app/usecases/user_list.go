/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package usecases

import "github.com/kazekim/UserData/app/entities"

func (u *defaultUserUseCase) ListUser() (*[]entities.User, error) {
	return u.userDBRepo.ListUser()
}

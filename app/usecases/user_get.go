/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package usecases

import "github.com/kazekim/UserData/app/entities"

func (u *defaultUserUseCase) GetUser(id int64) (*entities.User, error) {
	return u.userDBRepo.GetUser(id)
}

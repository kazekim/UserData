/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package usecases

func (u *defaultUserUseCase) DeleteUser(id int64) error {
	return u.userDBRepo.DeleteUser(id)
}

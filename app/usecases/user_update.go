/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package usecases

func (u *defaultUserUseCase) UpdateUser(id int64, email, fullName, telephone string) error {
	return u.userDBRepo.UpdateUser(id, email, fullName, telephone)
}

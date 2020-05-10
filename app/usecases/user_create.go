/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package usecases

func (u *defaultUserUseCase) CreateUser(email, fullName, telephone string) error {
	return u.userDBRepo.CreateUser(email, fullName, telephone)
}

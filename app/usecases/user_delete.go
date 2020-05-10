/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package usecases

func (u *defaultUserUseCase) DeleteUser(id int64) error {

	_, err := u.userDBRepo.GetUser(id)
	if err != nil {
		return err
	}

	return u.userDBRepo.DeleteUser(id)
}

/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package repositories

func (repo *defaultUserDBRepository) CreateUser(email, fullName, telephone string) error {
	return repo.client.CreateUserByField(email, fullName, telephone)
}

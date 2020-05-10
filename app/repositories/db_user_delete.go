/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package repositories

func (repo *defaultUserDBRepository) DeleteUser(id int64) error {
	return repo.client.DeleteUser(id)
}

/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package userdb

import userdbentities "github.com/kazekim/UserData/internal/userdb/v1/entities"

func (c *defaultClient) CreateUserByField(email, fullName, telephone string) error {
	statement, _ := c.db.Prepare("INSERT INTO " + userdbentities.TableUser + " (" + userdbentities.TableUserFieldEmail + ", " + userdbentities.TableUserFieldFullName +
		", " + userdbentities.TableUserFieldTelephone + ") VALUES (?, ?, ?)")
	_, err := statement.Exec(email, fullName, telephone)
	return err
}

func (c *defaultClient) CreateUser(user userdbentities.User) error {
	return c.CreateUserByField(user.Email, user.Fullname, user.Telephone)
}

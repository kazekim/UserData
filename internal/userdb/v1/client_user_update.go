/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package userdb

import (
	"fmt"
	"github.com/kazekim/UserData/internal/userdb/v1/entities"
)

func (c *defaultClient) UpdateUser(user *userdbentities.User) error {
	stmt, err := c.db.Prepare("update " + userdbentities.TableUser + " set " + userdbentities.TableUserFieldEmail + " = ?, " +
		userdbentities.TableUserFieldFullName + " = ?, " + userdbentities.TableUserFieldTelephone + " = ?, " +
		userdbentities.TableUserFieldUpdatedAt + " = CURRENT_TIMESTAMP where " +
		userdbentities.TableUserFieldID + " = ?")
	if err != nil {
		return err
	}
	fmt.Println(*user)
	_, err = stmt.Exec(user.Email, user.Fullname, user.Telephone, user.ID)
	if err != nil {
		return err
	}
	return nil
}
